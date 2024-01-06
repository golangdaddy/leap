
package models

import "net/http"

type OVERLAYS struct {
	Meta    Internals
	Fields FieldsOVERLAYS `json:"fields" firestore:"fields"`
}

func NewOVERLAYS(parent *Internals, fields FieldsOVERLAYS) *OVERLAYS {
	if parent == nil {
		return &OVERLAYS{
			Meta: (Internals{}).NewInternals("overlayss"),
			Fields: fields,
		}
	}
	return &OVERLAYS{
		Meta: parent.NewInternals("overlayss"),
		Fields: fields,
	}
}

type FieldsOVERLAYS struct {
	Name string `json:"name"`
	Type string `json:"type"`
	
}

func (x *OVERLAYS) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

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
