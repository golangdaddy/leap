package models

import "net/http"

type FONT struct {
	Meta    Internals
	Fields FieldsFONT `json:"fields" firestore:"fields"`
}

func (project *PROJECT) NewFONT(fields FieldsFONT) *FONT {
	return &FONT{
		Meta: project.Meta.NewInternals("fonts"),
		Fields: fields,
	}
}

type FieldsFONT struct {
	Name string `json:"name"`
	Description string `json:"description"`
	
}

func (x *FONT) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

	var exists bool
	
	x.Fields.Name, exists = AssertSTRING(w, m, "name")
	if !exists {
		return false
	}
	x.Fields.Description, exists = AssertSTRING(w, m, "description")
	if !exists {
		return false
	}
	return true
}
