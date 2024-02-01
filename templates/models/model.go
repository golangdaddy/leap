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
	{{if eq false .Options.Admin}}// this object inherits its admin permissions
	log.Println("OPTIONS ADMIN IS OFF:", parent.Moderation.Object)
	if len(parent.Moderation.Object) == 0 {
		log.Println("USING PARENT ID AS MODERATION OBJECT")
		object.Meta.Moderation.Object = parent.ID
	} else {
		log.Println("USING PARENT'S MODERATION OBJECT")
		object.Meta.Moderation.Object = parent.Moderation.Object
	}{{end}}
	{{if .Options.Admin}}// this object is owned by the user that created it
	log.Println("OPTIONS ADMIN IS ON:", user.Meta.ID)
	object.Meta.Moderation.Admins = append(
		object.Meta.Moderation.Admins,
		user.Meta.ID,
	){{end}}
	// add children to context
	object.Meta.Context.Children = []string{
		{{range .Children}}"{{.Name}}",{{end}}
	}
	return object
}

type Fields{{uppercase .Name}} struct {
	{{range .Fields}}{{titlecase .Name}} {{.Type}} `json:"{{lowercase .Name}}" firestore:"{{lowercase .Name}}"`
	{{end}}
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
	}

	x.Meta.Modify()

	return nil
}

{{end}}
