
package models

import "net/http"

type LAYER struct {
	Meta    Internals
	Fields FieldsLAYER `json:"fields" firestore:"fields"`
}

func NewLAYER(parent *Internals, fields FieldsLAYER) *LAYER {
	if parent == nil {
		return &LAYER{
			Meta: (Internals{}).NewInternals("layers"),
			Fields: fields,
		}
	}
	return &LAYER{
		Meta: parent.NewInternals("layers"),
		Fields: fields,
	}
}

type FieldsLAYER struct {
	Name string `json:"name"`
	Type string `json:"type"`
	
}

func (x *LAYER) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

	var exists bool
	
	x.Fields.Name, exists = AssertSTRING(w, m, "name")
	if !exists {
		return false
	}
	
	if !AssertRange(w, 1, 100, x.Fields.Name) {
		return false
	}
	x.Fields.Type, exists = AssertSTRING(w, m, "type")
	if !exists {
		return false
	}
	
	return true
}
