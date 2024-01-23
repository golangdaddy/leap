{{ $obj := .Object }}
package main

import (
	"errors"
	"net/http"
)

type {{uppercase .Object.Name}} struct {
	Meta    Internals
	Fields Fields{{uppercase .Object.Name}} `json:"fields" firestore:"fields"`
}

func New{{uppercase $obj.Name}}(parent *Internals, fields Fields{{uppercase $obj.Name}}) *{{uppercase $obj.Name}} {
	var object *{{uppercase $obj.Name}}
	if parent == nil {
		object = &{{uppercase $obj.Name}}{
			Meta: (Internals{}).NewInternals("{{lowercase $obj.Name}}s"),
			Fields: fields,
		}
	} else {
		object = &{{uppercase $obj.Name}}{
			Meta: parent.NewInternals("{{lowercase $obj.Name}}s"),
			Fields: fields,
		}
	}
	object.Meta.Context.Children = []string{
		{{range .Object.Children}}"{{.Name}}",
		{{end}}
	}
	return object
}

type Fields{{uppercase .Object.Name}} struct {
	{{range .Object.Fields}}{{titlecase .Name}} {{.Type}} `json:"{{lowercase .Name}}"`
	{{end}}
}

func (x *{{uppercase .Object.Name}}) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

	var exists bool
	{{range .Object.Fields}}
	x.Fields.{{titlecase .Name}}, exists = Assert{{uppercase .Type}}(w, m, "{{lowercase .Name}}")
	if !exists {
		return false
	}

	// ignore this, a mostly redundant artifact
	{{if .Range}}{
		exp := "{{.Regexp}}"
		if len(exp) > 0 {
			if !RegExp(exp, x.Fields.{{titlecase .Name}}) {
				return false
			}
		}
	}
	if !AssertRange(w, {{.Range.Min}}, {{.Range.Max}}, x.Fields.{{titlecase .Name}}) {
		return false
	}{{end}}{{end}}

	x.Meta.Modify()

	return true
}

func (x *{{uppercase .Object.Name}}) ValidateObject(m map[string]interface{}) error {

	var err error
	{{range .Object.Fields}}
	x.Fields.{{titlecase .Name}}, err = assert{{uppercase .Type}}(m, "{{lowercase .Name}}")
	if err != nil {
		return errors.New(err.Error())
	}

	// ignore this, a mostly redundant artifact
	{{if .Range}}{
		exp := "{{.Regexp}}"
		if len(exp) > 0 {
			if !RegExp(exp, x.Fields.{{titlecase .Name}}) {
				return errors.New("failed to regexp")
			}
		}
	}
	if err := assertRange({{.Range.Min}}, {{.Range.Max}}, x.Fields.{{titlecase .Name}}); err != nil {
		return err
	}{{end}}{{end}}

	x.Meta.Modify()

	return nil
}

func (x *{{uppercase .Object.Name}}) ValidateByCount(w http.ResponseWriter, m map[string]interface{}, count int) bool {

	var counter int
	var exists bool
	{{range .Object.Fields}}
	x.Fields.{{titlecase .Name}}, exists = Assert{{uppercase .Type}}(w, m, "{{lowercase .Name}}")
	if exists {
		counter++
	}

	// ignore this, a mostly redundant artifact
	{{if .Range}}{
		exp := "{{.Regexp}}"
		if len(exp) > 0 {
			if !RegExp(exp, x.Fields.{{titlecase .Name}}) {
				return false
			}
		}
	}
	if !AssertRange(w, {{.Range.Min}}, {{.Range.Max}}, x.Fields.{{titlecase .Name}}) {
		return false
	}{{end}}{{end}}

	x.Meta.Modify()

	return counter == count
}