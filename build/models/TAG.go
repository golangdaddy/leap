
package models

import (
	"net/http"
	"regexp"
)

type TAG struct {
	Meta    Internals
	Fields FieldsTAG `json:"fields" firestore:"fields"`
}

func NewTAG(parent *Internals, fields FieldsTAG) *TAG {
	if parent == nil {
		return &TAG{
			Meta: (Internals{}).NewInternals("tags"),
			Fields: fields,
		}
	}
	return &TAG{
		Meta: parent.NewInternals("tags"),
		Fields: fields,
	}
}

type FieldsTAG struct {
	Name string `json:"name"`
	Foreground_color string `json:"foreground_color"`
	Background_color string `json:"background_color"`
	
}

func (x *TAG) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

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
	if !AssertRange(w, 1, 10, x.Fields.Name) {
		return false
	}
	x.Fields.Foreground_color, exists = AssertSTRING(w, m, "foreground_color")
	if !exists {
		return false
	}
	
	{
		exp := ""
		if len(exp) > 0 {
			if !regexp.MustCompile(exp).MatchString(x.Fields.Foreground_color) {
				return false
			}
		}
	}
	if !AssertRange(w, 1, 10, x.Fields.Foreground_color) {
		return false
	}
	x.Fields.Background_color, exists = AssertSTRING(w, m, "background_color")
	if !exists {
		return false
	}
	
	{
		exp := ""
		if len(exp) > 0 {
			if !regexp.MustCompile(exp).MatchString(x.Fields.Background_color) {
				return false
			}
		}
	}
	if !AssertRange(w, 1, 10, x.Fields.Background_color) {
		return false
	}

	return true
}
