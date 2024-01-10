{{ $obj := .Object }}
package models

import (
	"net/http"
	"regexp"
)

type {{uppercase .Object.Name}} struct {
	Meta    Internals
	Fields Fields{{uppercase .Object.Name}} `json:"fields" firestore:"fields"`
}

func New{{uppercase $obj.Name}}(parent *Internals, fields Fields{{uppercase $obj.Name}}) *{{uppercase $obj.Name}} {
	if parent == nil {
		return &{{uppercase $obj.Name}}{
			Meta: (Internals{}).NewInternals("{{lowercase $obj.Name}}s"),
			Fields: fields,
		}
	}
	return &{{uppercase $obj.Name}}{
		Meta: parent.NewInternals("{{lowercase $obj.Name}}s"),
		Fields: fields,
	}
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
	{{if .Range}}
	{
		exp := "{{.Regexp}}"
		if len(exp) > 0 {
			if !regexp.MustCompile(exp).MatchString(x.Fields.{{titlecase .Name}}) {
				return false
			}
		}
	}
	if !AssertRange(w, {{.Range.Min}}, {{.Range.Max}}, x.Fields.{{titlecase .Name}}) {
		return false
	}{{end}}{{end}}

	return true
}
