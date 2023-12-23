package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

type Stack struct {
	Objects   []*Object `json:"objects"`
	objectMap map[string]*Object
}

func parseStack() (*Stack, error) {
	b, err := os.ReadFile("./projects/test/tree.json")
	if err != nil {
		return nil, err
	}
	stack := &Stack{}
	return stack, json.Unmarshal(b, stack)
}

func main() {

	os.RemoveAll("./build/")
	if err := os.Mkdir("./build/"); err != nil {
		panic(err)
	}

	stack, err := parseStack()
	if err != nil {
		log.Println(err)
		return
	}

	buf := bytes.NewBuffer(nil)

	for _, object := range stack.Objects {

		// Parse the template string
		model, err := loadTemplate("./templates/models/model.go")
		if err != nil {
			panic(err)
		}

		// Execute the template with the provided data
		err = model.Execute(buf, object)
		if err != nil {
			panic(err)
		}

		if err := os.WriteFile(
			fmt.Sprintf("./build/%s.go", strings.ToLower(object.Name)),
			buf.Bytes(),
			0777,
		); err != nil {
			panic(err)
		}
	}

}

func uppercase(s string) string {
	return strings.ToUpper(s)
}

func lowercase(s string) string {
	return strings.ToLower(s)
}

// loaTemplate Parses the template buffer
func loadTemplate(path string) (*template.Template, error) {

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	p := strings.Split(path, "/")
	templateName := p[len(p)-1]

	return template.New(templateName).Funcs(
		template.FuncMap{
			"lowercase": lowercase,
			"uppercase": uppercase,
		},
	).Parse(string(b))
}
