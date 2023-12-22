package main

import (
	"html/template"
	"os"
	"strings"
)

func main() {

	os.RemoveAll("./build/*")

	b, err := os.ReadFile("./templates/models/model.go")
	if err != nil {
		panic(err)
	}

	// Create a template and register the custom function
	tmpl := template.New("greeting").Funcs(
		template.FuncMap{
			"lowercase": lowercase,
			"uppercase": uppercase,
		},
	)

	// Parse the template string
	tmpl, err = tmpl.Parse(string(b))
	if err != nil {
		panic(err)
	}

	// Execute the template with the provided data
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}

func uppercase(s string) string {
	return strings.ToUpper(s)
}

func lowercase(s string) string {
	return strings.ToLower(s)
}
