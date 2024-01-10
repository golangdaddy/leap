
package models

import (
	"net/http"
	"regexp"
)

type OVERLAY struct {
	Meta    Internals
	Fields FieldsOVERLAY `json:"fields" firestore:"fields"`
}

func NewOVERLAY(parent *Internals, fields FieldsOVERLAY) *OVERLAY {
	if parent == nil {
		return &OVERLAY{
			Meta: (Internals{}).NewInternals("overlays"),
			Fields: fields,
		}
	}
	return &OVERLAY{
		Meta: parent.NewInternals("overlays"),
		Fields: fields,
	}
}

type FieldsOVERLAY struct {
	Name string `json:"name"`
	Type string `json:"type"`
	
}

func (x *OVERLAY) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

	var exists bool
	
	x.Fields.Name, exists = AssertSTRING(w, m, "name")
	if !exists {
		return false
	}
	
	{
		exp := ""
		if len(exp) > 0 {
			if !regexp.MustCompile(exp).MatchString(x.Fields.Name) {
				return false
			}
		}
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
