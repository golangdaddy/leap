
package models

import (
	"net/http"
	"regexp"
)

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
	Max_mint int `json:"max_mint"`
	
}

func (x *ELEMENT) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

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
	x.Fields.Max_mint, exists = AssertINT(w, m, "max_mint")
	if !exists {
		return false
	}
	

	return true
}
