package models

import (
	"encoding/json"
	"os"
)

type Config struct {
	WebAPI        string `json:"webAPI"`
	HostAPI       string `json:"hostAPI"`
	WebsocketHost string `json:"websocketHost"`
	RepoURI       string `json:"repoURI"`
	// project info
	ProjectID     string `json:"projectID"`
	ProjectName   string `json:"projectName"`
	ProjectRegion string `json:"projectRegion"`
}

type Stack struct {
	// the display name of the site
	WebsiteName string `json:"siteName"`
	Config      Config
	Objects     []*Object    `json:"objects"`
	Entrypoints []*Object    `json:"entrypoints"`
	Options     StackOptions `json:"options"`
}

type StackOptions struct {
	Sidebar             bool             `json:"sidebar"`
	ChatGPT             bool             `json:"chatgpt"`
	Assetlayer          bool             `json:"assetlayer"`
	Wallets             []string         `json:"wallets"`
	WhitelistDomains    bool             `json:"whitelistDomains"`
	RegistrationDomains []string         `json:"registrationDomains"`
	WhitelistEmails     bool             `json:"whitelistEmails"`
	RegistrationEmails  []string         `json:"registrationEmails"`
	Pusher              *OptionsPusher   `json:"pusher"`
	Handcash            *OptionsHandcash `json:"handcash"`
}

type OptionsHandcash struct {
	AppID     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type OptionsPusher struct {
	AppID   string `json:"appId"`
	Key     string `json:"key"`
	Secret  string `json:"secret"`
	Cluster string `json:"cluster"`
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
