
package models

import (
	"net/http"
	"regexp"
)

type FONT struct {
	Meta    Internals
	Fields FieldsFONT `json:"fields" firestore:"fields"`
}

func NewFONT(parent *Internals, fields FieldsFONT) *FONT {
	if parent == nil {
		return &FONT{
			Meta: (Internals{}).NewInternals("fonts"),
			Fields: fields,
		}
	}
	return &FONT{
		Meta: parent.NewInternals("fonts"),
		Fields: fields,
	}
}

type FieldsFONT struct {
	Name string `json:"name"`
	
}

func (x *FONT) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

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

	return true
}
