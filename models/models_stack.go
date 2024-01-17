package models

import (
	"encoding/json"
	"os"
)

type Stack struct {
	WebAPI      string    `json:"webAPI"`
	HostAPI     string    `json:"hostAPI"`
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
