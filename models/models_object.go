package models

type Object struct {
	Name        string         `json:"name"`
	Names       []string       `json:"names"`
	Plural      string         `json:"plural"`
	JSON        string         `json:"json"`
	Context     string         `json:"context"`
	Parents     []string       `json:"parents,omitempty"`
	ParentCount int            `json:"parentCount,omitempty"`
	Children    []*Object      `json:"children,omitempty"`
	Fields      []*Field       `json:"fields"`
	ListMode    string         `json:"listMode"`
	Options     Options        `json:"options"`
	Tags        []string       `json:"tags"`
	ChildTags   map[int]string `json:"childTags"`
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
	Member       *Member     `json:"member"`
	Job          bool        `json:"job"`
	Comment      bool        `json:"comment"`
	Order        bool        `json:"order"`
	File         bool        `json:"file"`
	Image        bool        `json:"image"`
	Photo        bool        `json:"photo"`
	EXIF         bool        `json:"exif"`
	Font         bool        `json:"font"`
	TopicCreate  *string     `json:"topicCreate"`
	Topics       []*JobTopic `json:"topics"`
	Assetlayer   *Assetlayer `json:"assetlayer"`
	Handcash     Handcash    `json:"handcash"`
	Pusher       bool        `json:"pusher"`
	Permissions  Permissions `json:"permissions"`
	FilterFields []*Field    `json:"filterFields"`
}

type Handcash struct {
	Type     string
	Payments []HandcashPayment
	Mint     []HandcashMint
}

type HandcashPayment struct {
	CurrencyCode string
	To           string
	Amount       float64
}

type HandcashMint struct {
	Data map[string]interface{}
}

type Member struct {
	View bool `json:"view"`
	Edit bool `json:"edit"`
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
