package leap

import (
	"encoding/json"
	"strings"
	"text/template"

	"github.com/golangdaddy/leap/models"
)

var funcMap = template.FuncMap{
	"parentcount":      parentcount,
	"lowercase":        lowercase,
	"uppercase":        uppercase,
	"titlecase":        titlecase,
	"json":             jsonmarshal,
	"tidy":             tidy,
	"firstparent":      firstparent,
	"firstparenttitle": firstparenttitle,
	"stringslength":    stringslength,
	"jsonmarshal":      jsonmarshal,
}

func stringslength(a []string) int {
	return len(a)
}

func firstparent(a []string) string {
	return a[0]
}

func firstparenttitle(a []string) string {
	return titlecase(a[0])
}

func jsonmarshal(x interface{}) string {
	b, _ := json.Marshal(x)
	return string(b)
}

func titlecase(s string) string {
	return string(strings.ToUpper(s)[0]) + string(strings.ToLower(s)[1:])
}

func uppercase(s string) string {
	return strings.ToUpper(s)
}

func lowercase(s string) string {
	return strings.ToLower(s)
}

func tidy(s string) string {
	s = strings.Replace(s, "[]", "array", -1)
	s = strings.Replace(s, "[", "", -1)
	s = strings.Replace(s, "]", "", -1)
	s = strings.Replace(s, "{", "", -1)
	s = strings.Replace(s, "}", "", -1)
	return strings.ToUpper(s)
}

func parentcount(object *models.Object) int {
	return object.ParentCount * 2
}
