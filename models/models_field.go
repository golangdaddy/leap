package models

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
	RegexpHex    string   `json:"regexpHex"`
}

func (object *Object) NewField() *Field {
	return &Field{}
}

type Range struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}
