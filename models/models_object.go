package models

type Object struct {
	Name        string    `json:"name"`
	Names       []string  `json:"names"`
	Plural      string    `json:"plural"`
	JSON        string    `json:"json"`
	Mode        string    `json:"mode"`
	Context     string    `json:"context"`
	Parents     []string  `json:"parents,omitempty"`
	ParentCount int       `json:"parentCount,omitempty"`
	Children    []*Object `json:"children,omitempty"`
	Fields      []*Field  `json:"fields"`
	ListMode    string    `json:"listMode"`
	Options     Options   `json:"options"`
}

const (
	ListModeCreated  = "created"
	ListModeModified = "modified"
	ListModeOrder    = "ordered"
	ListModeTimeline = "exif"
)

type Options struct {
	ReadOnly     bool        `json:"readonly"`
	Admin        bool        `json:"admin"`
	Job          bool        `json:"job"`
	Order        bool        `json:"order"`
	File         bool        `json:"file"`
	Image        bool        `json:"image"`
	Photo        bool        `json:"photo"`
	EXIF         bool        `json:"exif"`
	Font         bool        `json:"font"`
	TopicCreate  *string     `json:"topicCreate"`
	Topics       []*JobTopic `json:"topics"`
	Assetlayer   *Assetlayer `json:"assetlayer"`
	Pusher       bool        `json:"pusher"`
	Permissions  Permissions `json:"permissions"`
	FilterFields []*Field    `json:"filterFields"`
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
