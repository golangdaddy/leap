package main

import (
	"encoding/json"
	"strings"
	"text/template"

	"github.com/golangdaddy/leap/models"
)

var funcMap = template.FuncMap{
	"parentcount": parentcount,
	"lowercase":   lowercase,
	"uppercase":   uppercase,
	"titlecase":   titlecase,
	"json":        jsonmarshal,
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

func parentcount(object *models.Object) int {
	return object.ParentCount * 2
}
