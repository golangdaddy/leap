package main

type Object struct {
	ParentName string   `json:"parentName"`
	Name       string   `json:"name"`
	Fields     []*Field `json:"fields"`
}

type Field struct {
	Name string `json:"name"`
	// go primative types
	Type string `json:"type"`
	// define frontend options
	Meta map[string]string `json:"meta"`
}
