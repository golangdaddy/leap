package models

type Object struct {
	JSON        string    `json:"json"`
	Mode        string    `json:"mode"`
	Context     string    `json:"context"`
	Parents     []string  `json:"parents,omitempty"`
	ParentCount int       `json:"parentCount,omitempty"`
	Children    []*Object `json:"children,omitempty"`
	Name        string    `json:"name"`
	Fields      []*Field  `json:"fields"`
	Options     Options   `json:"options"`
}

type Options struct {
	Order          bool        `json:"order"`
	File           bool        `json:"file"`
	Image          bool        `json:"image"`
	Font           bool        `json:"font"`
	UseCreateTopic bool        `json:"useCreateTopic"`
	TopicCreate    string      `json:"topicCreate"`
	Assetlayer     Assetlayer  `json:"assetlayer"`
	Permissions    Permissions `json:"permissions"`
}

type Permissions struct {
	AdminsOnly bool
	AdminsEdit bool
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
	Context string `json:"context"`
	JSON    string `json:"json"`
	Name    string `json:"name"`
	// go primative types
	Type string `json:"type"`
	// define frontend options
	Input        string   `json:"input"`
	InputOptions []string `json:"inputOptions,omitempty"`
	Required     bool     `json:"required"`
	Range        *Range   `json:"range"`
	Regexp       string   `json:"regexp"`
}

type Range struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

func (object *Object) NewField() *Field {
	return &Field{}
}
