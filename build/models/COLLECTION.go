
package models

import (
	"net/http"
	"regexp"
)

type COLLECTION struct {
	Meta    Internals
	Fields FieldsCOLLECTION `json:"fields" firestore:"fields"`
}

func NewCOLLECTION(parent *Internals, fields FieldsCOLLECTION) *COLLECTION {
	if parent == nil {
		return &COLLECTION{
			Meta: (Internals{}).NewInternals("collections"),
			Fields: fields,
		}
	}
	return &COLLECTION{
		Meta: parent.NewInternals("collections"),
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
	x.Fields.Description, exists = AssertSTRING(w, m, "description")
	if !exists {
		return false
	}
	
	{
		exp := ""
		if len(exp) > 0 {
			if !regexp.MustCompile(exp).MatchString(x.Fields.Description) {
				return false
			}
		}
	}
	if !AssertRange(w, 1, 100, x.Fields.Description) {
		return false
	}

	return true
}
