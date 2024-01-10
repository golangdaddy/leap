
package models

import (
	"net/http"
	"regexp"
)

type PROJECT struct {
	Meta    Internals
	Fields FieldsPROJECT `json:"fields" firestore:"fields"`
}

func NewPROJECT(parent *Internals, fields FieldsPROJECT) *PROJECT {
	if parent == nil {
		return &PROJECT{
			Meta: (Internals{}).NewInternals("projects"),
			Fields: fields,
		}
	}
	return &PROJECT{
		Meta: parent.NewInternals("projects"),
		Fields: fields,
	}
}

type FieldsPROJECT struct {
	Name string `json:"name"`
	Description string `json:"description"`
	
}

func (x *PROJECT) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

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
