package models

import (
	"encoding/json"
	"os"
)

type Stack struct {
	WebAPI        string       `json:"webAPI"`
	HostAPI       string       `json:"hostAPI"`
	WebsocketHost string       `json:"websocketHost"`
	RepoURI       string       `json:"repoURI"`
	SiteName      string       `json:"siteName"`
	ProjectID     string       `json:"projectID"`
	DatabaseID    string       `json:"databaseID"`
	Objects       []*Object    `json:"objects"`
	Entrypoints   []string     `json:"entrypoints"`
	Options       StackOptions `json:"options"`
}

type StackOptions struct {
	Pusher              bool     `json:"pusher"`
	Assetlayer          bool     `json:"assetlayer"`
	Wallets             []string `json:"wallets"`
	WhitelistDomains    bool     `json:"whitelistDomains"`
	RegistrationDomains []string `json:"registrationDomains"`
	WhitelistEmails     bool     `json:"whitelistEmails"`
	RegistrationEmails  []string `json:"registrationEmails"`
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
