package models

import "net/http"

type COLLECTION struct {
	Meta    Internals
	Fields FieldsCOLLECTION `json:"fields" firestore:"fields"`
}

func (project *PROJECT) NewCOLLECTION(fields FieldsCOLLECTION) *COLLECTION {
	return &COLLECTION{
		Meta: project.Meta.NewInternals("collections"),
		Fields: fields,
	}
}

type FieldsCOLLECTION struct {
	Name string `json:"name"`
	Description string `json:"description"`
	
}

func (x *COLLECTION) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

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
