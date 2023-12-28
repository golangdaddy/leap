package models

import "net/http"

type LAYER struct {
	Meta    Internals
	Fields FieldsLAYER `json:"fields" firestore:"fields"`
}

func (collection *COLLECTION) NewLAYER(fields FieldsLAYER) *LAYER {
	return &LAYER{
		Meta: collection.Meta.NewInternals("layers"),
		Fields: fields,
	}
}

type FieldsLAYER struct {
	Name string `json:"name"`
	Description string `json:"description"`
	
}

func (x *LAYER) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

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
