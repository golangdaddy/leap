package models

import (
	"encoding/json"
	"os"
)

type Stack struct {
	RepoURI     string    `json:"repoURI"`
	SiteName    string    `json:"siteName"`
	ProjectID   string    `json:"projectID"`
	DatabaseID  string    `json:"databaseID"`
	Objects     []*Object `json:"objects"`
	Wallets     []string  `json:"wallets"`
	Entrypoints []string  `json:"entrypoints"`
}

func (stack *Stack) NewObject(parent *Object, name string) *Object {
	obj := &Object{
		Name: name,
	}
	if parent != nil {
		obj.Parents = append(obj.Parents, parent.Name)
	}
	stack.Objects = append(stack.Objects, obj)
	return obj
}

func ParseStack(folder, file string) (*Stack, error) {
	path := folder + "/" + file + ".json"
	println("parsing path", path)
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	stack := &Stack{}
	return stack, json.Unmarshal(b, stack)
}

type Object struct {
	JSON string `json:"json"`
	Mode string `json:"mode"`

	Parents     []string   `json:"parents,omitempty"`
	ParentCount int        `json:"parentCount,omitempty"`
	Children    []*Object  `json:"children,omitempty"`
	Name        string     `json:"name"`
	Fields      []*Field   `json:"fields"`
	Options     Options    `json:"options"`
	Assetlayer  Assetlayer `json:"assetlayer"`
}

type Options struct {
	Order          bool   `json:"order"`
	File           bool   `json:"file"`
	Image          bool   `json:"image"`
	Font           bool   `json:"font"`
	UseCreateTopic bool   `json:"useCreateTopic"`
	TopicCreate    string `json:"topicCreate"`
}

type Assetlayer struct {
	Token  bool `json:"token"`
	Wallet bool `json:"wallet"`
}

type ObjectRef struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

func (object *Object) HasParent() bool {
	return len(object.Parents) > 0
}

type Field struct {
	JSON string `json:"json"`
	Name string `json:"name"`
	// go primative types
	Type string `json:"type"`
	// define frontend options
	Input        string   `json:"input"`
	InputOptions []string `json:"inputOptions,omitempty"`
	Required     bool     `json:"required"`
	Range        *Range   `json:"range"`
	Regexp       string   `json:"regexp`
}

type Range struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

func (object *Object) NewField() *Field {
	return &Field{}
}
