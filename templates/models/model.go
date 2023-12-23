package models

import "net/http"

type {{uppercase .Name}} struct {
	Meta    Internals
	Fields Fields{{uppercase .Name}} `json:"fields" firestore:"fields"`
}

func ({{lowercase .ParentName}} *{{.ParentName}}) New{{.Name}}(name string, fields *Fields{{uppercase .Name}}) *{{.Name}} {
	return &{{.Name}}{
		Meta: {{lowercase .ParentName}}.Meta.NewInternals("{{lowercase .Name}}s"),
		Name: name,
		Fields: fields,
	}
}

type Fields{{uppercase .Name}} struct {
	{{range .Fields}}{{.Name}} {{.Type}}
	{{end}}
}

func ({{lowercase .Name}} *{{uppercase .Name}}) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

	var exists bool
	{{range .Fields}}
	{{lowercase .Name}}.Name, exists = Assert{{uppercase .Type}}(w, m, "name")
	if !exists {
		return false
	}{{end}}
	return true
}
