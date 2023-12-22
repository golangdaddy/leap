package models

import "net/http"

type {{.ModelName}} struct {
	Meta    Internals
	Fields {{.ModelName}}Options `json:"options" firestore:"options"`
}

func ({{lowercase .ParentModelName}} *{{.ParentModelName}}) New{{.ModelName}}(name string, fields *{{.ModelName}}Fields) *{{.ModelName}} {
	return &{{.ModelName}}{
		Meta: project.Meta.NewInternals("{{lowercase .ModelName}}s"),
		Name: name,
		Fields: fields,
	}
}

type {{.ModelName}} Fields struct {
	{{range .Fields}}
	{{.Name}}: {{.Type}},
	{{end}}
}

func ({{lowercase .ModelName}} *{{.ModelName}}) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

	var exists bool

	{{range .Fields}}
	{{lowercase .ModelName}}.Name, exists = AssertKeyValue(w, m, "name")
	if !exists {
		return false
	}
	{{end}}

	return true
}
