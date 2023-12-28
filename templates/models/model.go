package models

import "net/http"

type {{uppercase .Object.Name}} struct {
	Meta    Internals
	Fields Fields{{uppercase .Object.Name}} `json:"fields" firestore:"fields"`
}

func ({{lowercase .Object.ParentName}} *{{uppercase .Object.ParentName}}) New{{uppercase .Object.Name}}(fields Fields{{uppercase .Object.Name}}) *{{uppercase .Object.Name}} {
	return &{{uppercase .Object.Name}}{
		Meta: {{lowercase .Object.ParentName}}.Meta.NewInternals("{{lowercase .Object.Name}}s"),
		Fields: fields,
	}
}

type Fields{{uppercase .Object.Name}} struct {
	{{range .Object.Fields}}{{.Name}} {{.Type}} `json:"{{lowercase .Name}}"`
	{{end}}
}

func (x *{{uppercase .Object.Name}}) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

	var exists bool
	{{range .Object.Fields}}
	x.Fields.{{.Name}}, exists = Assert{{uppercase .Type}}(w, m, "{{lowercase .Name}}")
	if !exists {
		return false
	}{{end}}
	return true
}
