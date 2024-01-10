
package models

import (
	"net/http"
	"regexp"
)

type ATTRIBUTE struct {
	Meta    Internals
	Fields FieldsATTRIBUTE `json:"fields" firestore:"fields"`
}

func NewATTRIBUTE(parent *Internals, fields FieldsATTRIBUTE) *ATTRIBUTE {
	if parent == nil {
		return &ATTRIBUTE{
			Meta: (Internals{}).NewInternals("attributes"),
			Fields: fields,
		}
	}
	return &ATTRIBUTE{
		Meta: parent.NewInternals("attributes"),
		Fields: fields,
	}
}

type FieldsATTRIBUTE struct {
	Name string `json:"name"`
	Min int `json:"min"`
	Max int `json:"max"`
	
}

func (x *ATTRIBUTE) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

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
	x.Fields.Min, exists = AssertINT(w, m, "min")
	if !exists {
		return false
	}
	
	x.Fields.Max, exists = AssertINT(w, m, "max")
	if !exists {
		return false
	}
	

	return true
}
