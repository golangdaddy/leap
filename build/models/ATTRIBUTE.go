package models

import "net/http"

type ATTRIBUTE struct {
	Meta    Internals
	Fields FieldsATTRIBUTE `json:"fields" firestore:"fields"`
}

func (collection *COLLECTION) NewATTRIBUTE(fields FieldsATTRIBUTE) *ATTRIBUTE {
	return &ATTRIBUTE{
		Meta: collection.Meta.NewInternals("attributes"),
		Fields: fields,
	}
}

type FieldsATTRIBUTE struct {
	Name string `json:"name"`
	Description string `json:"description"`
	
}

func (x *ATTRIBUTE) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

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
