
package models

import "net/http"

type ELEMENT struct {
	Meta    Internals
	Fields FieldsELEMENT `json:"fields" firestore:"fields"`
}

func NewELEMENT(parent *Internals, fields FieldsELEMENT) *ELEMENT {
	if parent == nil {
		return &ELEMENT{
			Meta: (Internals{}).NewInternals("elements"),
			Fields: fields,
		}
	}
	return &ELEMENT{
		Meta: parent.NewInternals("elements"),
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
	
	if !AssertRange(w, 1, 100, x.Fields.Name) {
		return false
	}
	x.Fields.Description, exists = AssertSTRING(w, m, "description")
	if !exists {
		return false
	}
	
	if !AssertRange(w, 1, 100, x.Fields.Description) {
		return false
	}
	return true
}
