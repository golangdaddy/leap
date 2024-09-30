{{range .Objects}}

type {{uppercase .Name}} struct {
	Meta    Internals
	Fields Fields{{uppercase .Name}} `json:"fields" firestore:"fields"`
}

func (user *User) New{{uppercase .Name}}(parent *Internals, fields Fields{{uppercase .Name}}) *{{uppercase .Name}} {
	var object *{{uppercase .Name}}
	if parent == nil {
		object = &{{uppercase .Name}}{
			Meta: (Internals{}).NewInternals("{{lowercase .Name}}s"),
			Fields: fields,
		}
	} else {
		object = &{{uppercase .Name}}{
			Meta: parent.NewInternals("{{lowercase .Name}}s"),
			Fields: fields,
		}
	}

	object.Meta.ClassName = "{{lowercase .Plural}}"
	object.Meta.Context.User = user.Meta.ID

	colors, err := gamut.Generate(8, gamut.PastelGenerator{})
	if err != nil {
		log.Println(err)
	} else {
		object.Meta.Media.Color = gamut.ToHex(colors[0])
	}

	{{if eq false .Options.Admin}}// this object inherits its admin permissions
		if parent != nil {
			log.Println("OPTIONS ADMIN IS OFF:", parent.Moderation.Object)
			if len(parent.Moderation.Object) == 0 {
				log.Println("USING PARENT ID AS MODERATION OBJECT")
				object.Meta.Moderation.Object = parent.ID
			} else {
				log.Println("USING PARENT'S MODERATION OBJECT")
				object.Meta.Moderation.Object = parent.Moderation.Object
			}
		}
	{{end}}

	{{if .Options.Admin}}
		// this object is owned by the user that created it
		log.Println("OPTIONS ADMIN IS ON:", user.Meta.ID)
		object.Meta.Moderation.Admins = append(
			object.Meta.Moderation.Admins,
			user.Meta.ID,
		)
	{{end}}

	{{if .Options.Handcash}}
	{{if eq "pay" .Options.Handcash.Type}}
	{{range .Options.Handcash.Payments}}
	object.Meta.Payment.Destinations = append(
		object.Meta.Payment.Destinations,
		&PaymentDestination{
			To: "{{.To}}",
			CurrencyCode: "{{.CurrencyCode}}",
			Amount: {{.Amount}},
		},
	)
	{{end}}{{end}}{{end}}

	// add children to context
	object.Meta.Context.Children = []string{
		{{range .Children}}"{{.Name}}",{{end}}
	}
	return object
}

// set the fields export tags to lowercase
type Fields{{uppercase .Name}} struct {
	{{range $i, $field := .Fields}}
		{{if eq nil $field.Element}}
			{{range $index, $input := $field.Inputs}}
				{{if ne nil $input.Element}}
					{{$input.ID}} {{$input.Element.Go}} `json:"{{lowercase $field.ID}}" firestore:"{{lowercase $field.ID}}"`
				{{end}}
			{{end}}
		{{else}}
			{{$field.ID}} {{$field.Element.Go}} `json:"{{lowercase $field.ID}}" firestore:"{{lowercase $field.ID}}"`
		{{end}}
	{{end}}
}

func (x *{{uppercase .Name}}) Schema() *models.Object {
	obj := &models.Object{}
	json.Unmarshal([]byte(`{{jsonmarshal .}}`), obj)
	return obj
}

func (x *{{uppercase .Name}}) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {
	if err := x.ValidateObject(m); err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return false
	}
	return true
}

func (x *{{uppercase .Name}}) ValidateObject(m map[string]interface{}) error {

	var err error
	var exists bool
	{{range $i, $field := .Fields}}

	_, exists = m["{{$field.ID}}"]
	if {{.Required}} && !exists {
		return errors.New("required field '{{$field.ID}}' not supplied")
	}
	if exists {
		var exp string
		{{if eq nil $field.Element}}
			{{range $index, $subfield := $field.Inputs}}
				x.Fields.{{$subfield.ID}}, err = assert{{uppercase $subfield.Element.Go}}(m, "{{.ID}}")
				if err != nil {
					return errors.New(err.Error())
				}
				exp = "{{.RegexpHex}}"
				if len(exp) > 0 {
					log.Println("EXPR", exp)
					b, err := hex.DecodeString(exp)
					if err != nil {
						log.Println(err)
					}
					if !RegExp(string(b), fmt.Sprintf("%v", x.Fields.{{.ID}})) {
						return fmt.Errorf("failed to regexpHex: %s >> %s", string(b), x.Fields.{{.ID}})
					}
				}
				{{if .Range}}
					if err := assertRangeMin({{.Range.Min}}, x.Fields.{{.ID}}); err != nil {
						{{if .Required}}
						return err
						{{end}}
					}
					if err := assertRangeMax({{.Range.Max}}, x.Fields.{{.ID}}); err != nil {
						return err
					}
				{{end}}
			{{end}}
		{{else}}
			x.Fields.{{$field.ID}}, err = assert{{uppercase $field.Element.Go}}(m, "{{$field.ID}}")
			if err != nil {
				return errors.New(err.Error())
			}
			exp = "{{.RegexpHex}}"
			if len(exp) > 0 {
				log.Println("EXPR", exp)
				b, err := hex.DecodeString(exp)
				if err != nil {
					log.Println(err)
				}
				if !RegExp(string(b), fmt.Sprintf("%v", x.Fields.{{.ID}})) {
					return fmt.Errorf("failed to regexpHex: %s >> %s", string(b), x.Fields.{{.ID}})
				}
			}
			{{if .Range}}
				if err := assertRangeMin({{.Range.Min}}, x.Fields.{{.ID}}); err != nil {
					{{if .Required}}
					return err
					{{end}}
				}
				if err := assertRangeMax({{.Range.Max}}, x.Fields.{{.ID}}); err != nil {
					return err
				}
			{{end}}
		{{end}}
	}
	{{end}}

	// extract name field if exists
	name, ok := m["name"].(string)
	if ok {
		x.Meta.Name = name	
	} else {
		log.Println("trying to composite object name")
		var names []string
		{{range .Names}}
			names = append(names, m["{{.}}"].(string))
		{{end}}
		x.Meta.Name = strings.Join(names, " ")
	}

	x.Meta.Modify()

	return nil
}

// assert file is an image because of .Object.Options.Image
func (object *{{uppercase .Name}}) ValidateImage{{uppercase .Name}}(fileBytes []byte) (image.Image, error) {

	img, _, err := image.Decode(bytes.NewBuffer(fileBytes))
	if err != nil {
		return nil, err
	}
	object.Meta.Media.Image = true

	// determine image format
	if jpegstructure.NewJpegMediaParser().LooksLikeFormat(fileBytes) {
		object.Meta.Media.Format = "JPEG"
	} else {
		if pngstructure.NewPngMediaParser().LooksLikeFormat(fileBytes) {
			object.Meta.Media.Format = "PNG"
		}
	}

	// Parse the EXIF data
	exifData, err := exif.Decode(bytes.NewBuffer(fileBytes))
	if err == nil {
		println(exifData.String())
		
		object.Meta.Media.EXIF = map[string]interface{}{}
	
		tm, err := exifData.DateTime()
		if err == nil {
			object.Meta.Media.EXIF["taken"] = tm.UTC().Unix()
			object.Meta.Modified = tm.UTC().Unix()
			fmt.Println("Taken: ", tm)
		}
	
		lat, long, err := exifData.LatLong()
		if err != nil {
			object.Meta.Media.EXIF["lat"] = lat
			object.Meta.Media.EXIF["lng"] = long
			fmt.Println("lat, long: ", lat, ", ", long)
		}
	}

	return img, nil
}

{{end}}
