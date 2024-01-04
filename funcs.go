package main

import (
	"encoding/json"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"lowercase": lowercase,
	"uppercase": uppercase,
	"titlecase": titlecase,
	"json":      jsonmarshal,
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
