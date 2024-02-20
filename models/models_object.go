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
	ReadOnly    bool        `json:"readonly"`
	Admin       bool        `json:"admin"`
	Job         bool        `json:"job"`
	Order       bool        `json:"order"`
	Color       bool        `json:"color"`
	File        bool        `json:"file"`
	Image       bool        `json:"image"`
	Font        bool        `json:"font"`
	TopicCreate *string     `json:"topicCreate"`
	Topics      []*JobTopic `json:"topics"`
	Assetlayer  *Assetlayer `json:"assetlayer"`
	Pusher      bool        `json:"pusher"`
	Permissions Permissions `json:"permissions"`
}

type JobTopic struct {
	Name  string `json:"name"`
	Topic string `json:"topic"`
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
