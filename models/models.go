package models

import (
	"encoding/json"
	"os"
)

type Stack struct {
	ProjectID  string    `json:"projectID"`
	DatabaseID string    `json:"databaseID"`
	Objects    []*Object `json:"objects"`
	objectMap  map[string]*Object
}

func (stack *Stack) NewObject(parent *Object, name string) *Object {
	obj := &Object{
		Name: name,
	}
	if parent != nil {
		obj.ParentName = parent.Name
	}
	stack.Objects = append(stack.Objects, obj)
	return obj
}

func ParseStack() (*Stack, error) {
	b, err := os.ReadFile("./projects/test/tree.json")
	if err != nil {
		return nil, err
	}
	stack := &Stack{}
	return stack, json.Unmarshal(b, stack)
}

type Object struct {
	ParentName string       `json:"parentName"`
	Name       string       `json:"name"`
	Fields     []*Field     `json:"fields"`
	Children   []*ObjectRef `json:"children"`
}

type ObjectRef struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

func (object *Object) NewChild(name string, priority int) {
	object.Children = append(object.Children, &ObjectRef{name, priority})
}

func (object *Object) HasParent() bool {
	return len(object.ParentName) > 0
}

func (object *Object) HasChildren() bool {
	return len(object.Children) > 0
}

type Field struct {
	Name string `json:"name"`
	// go primative types
	Type string `json:"type"`
	// define frontend options
	Input           string   `json:"input"`
	Required        bool     `json:"required"`
	Options         []string `json:"options"`
	Collection      string   `json:"collection"`
	CollectionField string   `json:"collectionField"`
}

func (object *Object) NewField() *Field {
	return &Field{}
}
