{{range .Objects}}

type {{.ID}} struct {
	Meta    Internals
	Fields Fields{{.ID}} `json:"fields" firestore:"fields"`
}

func (user *User) New{{.ID}}(parent *Internals, fields Fields{{.ID}}) *{{.ID}} {
	var object *{{.ID}}
	if parent == nil {
		object = &{{.ID}}{
			Meta: (Internals{}).NewInternals("{{lowercase .Name}}s"),
			Fields: fields,
		}
	} else {
		object = &{{.ID}}{
			Meta: parent.NewInternals("{{.ID}}s"),
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
	}{{end}}

	{{if .Options.Admin}}// this object is owned by the user that created it
	log.Println("OPTIONS ADMIN IS ON:", user.Meta.ID)
	object.Meta.Moderation.Admins = append(
		object.Meta.Moderation.Admins,
		user.Meta.ID,
	){{end}}

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

type Fields{{.ID}} struct {
	{{range .Fields}}{{titlecase .Name}} {{.Type}} `json:"{{.ID}}" firestore:"{{.ID}}"`
	{{end}}
}

func (x *{{.ID}}) Schema() *models.Object {
	obj := &models.Object{}
	json.Unmarshal([]byte(`{{jsonmarshal .}}`), obj)
	return obj
}

func (x *{{.ID}}) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {
	if err := x.ValidateObject(m); err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return false
	}
	return true
}

func (x *{{.ID}}) ValidateObject(m map[string]interface{}) error {

	var err error
	var exists bool
	{{range .Fields}}

	_, exists = m["{{lowercase .Name}}"]
	if {{.Required}} && !exists {
		return errors.New("required field '{{lowercase .Name}}' not supplied")
	}
	if exists {
		x.Fields.{{titlecase .Name}}, err = assert{{tidy .Type}}(m, "{{lowercase .Name}}")
		if err != nil {
			return errors.New(err.Error())
		}
		{
			exp := "{{.Regexp}}"
			if len(exp) > 0 {
				if !RegExp(exp, fmt.Sprintf("%v", x.Fields.{{titlecase .Name}})) {
					return fmt.Errorf("failed to regexp: %s >> %s", exp, x.Fields.{{titlecase .Name}})
				}
			}
		}
		{
			exp := "{{.RegexpHex}}"
			if len(exp) > 0 {
				log.Println("EXPR", exp)
				b, err := hex.DecodeString(exp)
				if err != nil {
					log.Println(err)
				}
				if !RegExp(string(b), fmt.Sprintf("%v", x.Fields.{{titlecase .Name}})) {
					return fmt.Errorf("failed to regexpHex: %s >> %s", string(b), x.Fields.{{titlecase .Name}})
				}
			}
		}
		{{if .Range}}
		if err := assertRangeMin({{.Range.Min}}, x.Fields.{{titlecase .Name}}); err != nil {
			{{if .Required}}
			return err
			{{end}}
		}
		if err := assertRangeMax({{.Range.Max}}, x.Fields.{{titlecase .Name}}); err != nil {
			return err
		}
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
		{{range .Names}}names = append(names, m["{{.}}"].(string))
		{{end}}
		x.Meta.Name = strings.Join(names, " ")
	}

	x.Meta.Modify()

	return nil
}

// assert file is an image because of .Object.Options.Image
func (object *{{.ID}}) ValidateImage{{.ID}}(fileBytes []byte) (image.Image, error) {

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
