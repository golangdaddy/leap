package models

import "net/http"

type PROJECT struct {
	Meta    Internals
	Fields FieldsPROJECT `json:"fields" firestore:"fields"`
}

func NewPROJECT(fields FieldsPROJECT) *PROJECT {
	return &PROJECT{
		Meta: (Internals{}).NewInternals("projects"),
		Fields: fields,
	}
}

type FieldsPROJECT struct {
	Name string `json:"name"`
	Description string `json:"description"`
	
}

func (x *PROJECT) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

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
