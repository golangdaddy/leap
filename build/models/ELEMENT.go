package models

import "net/http"

type ELEMENT struct {
	Meta    Internals
	Fields FieldsELEMENT `json:"fields" firestore:"fields"`
}

func (layer *LAYER) NewELEMENT(fields FieldsELEMENT) *ELEMENT {
	return &ELEMENT{
		Meta: layer.Meta.NewInternals("elements"),
		Fields: fields,
	}
}

type FieldsELEMENT struct {
	Name string `json:"name"`
	Description string `json:"description"`
	
}

func (x *ELEMENT) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

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
