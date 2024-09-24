package models

import (
	"encoding/hex"
	"strconv"
)

const (
	STRING   = "input"
	TEXT     = "textarea"
	NUMBER   = "number"
	DATE     = "date"
	PHONE    = "phone"
	EMAIL    = "email"
	CHECKBOX = "checkbox"
	SELECT   = "select"
)

type Field struct {
	Context string `json:"context"`
	Name    string `json:"name"`
	// go primative types
	Type string `json:"type"`
	// define frontend options
	Element        string   `json:"element"`
	Inputs         []*Field `json:"inputs,omitempty"`
	InputReference string   `json:"inputReference"`
	InputOptions   []string `json:"inputOptions,omitempty"`
	Required       bool     `json:"required"`
	Filter         bool     `json:"filter"`
	Range          *Range   `json:"range,omitempty"`
	Regexp         string   `json:"regexp"`
	RegexpHex      string   `json:"regexpHex"`
}

func (object *Object) NewField() *Field {
	return &Field{}
}

type Range struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

func (f *Field) SetName(s string) *Field {
	f.Name = s
	return f
}

func (f *Field) SetCtx(s string) *Field {
	f.Context = s
	return f
}

func Required(name string, args ...string) (f *Field) {
	f = Get(name, args...)
	f.Required = true
	return
}

func Get(name string, args ...string) (f *Field) {

	f = &Field{
		Type: name,
	}

	switch f.Type {
	case "address":
		f.Context = "An address of a location"
		f.Inputs = []*Field{
			Required("int").SetName("building number"),
			Get("int").SetName("apartment number"),
			Required("string", "75").SetName("street"),
			Required("string", "50").SetName("town or city"),
			Required("string", "50").SetName("country"),
		}
	case "float64":
		f.Context = "64-bit floating-point number"
		f.Element = NUMBER
		switch len(args) {
		case 0:
		case 1:
			x, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				panic(err)
			}
			f.Range = &Range{}
			f.Range.Max = x
		case 2:
			x, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				panic(err)
			}
			y, err := strconv.ParseFloat(args[1], 64)
			if err != nil {
				panic(err)
			}
			f.Range = &Range{}
			f.Range.Min = x
			f.Range.Max = y
		default:
			panic("invalid arg length: " + f.Type)
		}
	case "uint":
		f.Range = &Range{}
		f.Range.Min = 0.0
	case "int":
		f.Context = "any integer"
		f.Element = NUMBER
		switch len(args) {
		case 0:
		case 1:
			x, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			f.Range = &Range{}
			f.Range.Max = float64(x)
		case 2:
			x, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(args[1])
			if err != nil {
				panic(err)
			}
			f.Range = &Range{}
			f.Range.Min = float64(x)
			f.Range.Max = float64(y)
		default:
			panic("invalid arg length: " + f.Type)
		}
	case "date":
		f.Context = "A date either past or future"
		f.Element = DATE
	case "email":
		f.Context = "An email address"
		f.Element = EMAIL
	case "name":
		f.Context = "A name or names of something or someone"
		f.Element = STRING
		f.Range = &Range{}
		f.Range.Min = 2
		f.Range.Max = 50
	case "name.person", "person.name":
		f.Context = "A full legal name of a person"
		f.Inputs = []*Field{
			Required("name").SetName("first-name"),
			Get("name").SetName("middle-names"),
			Required("name").SetName("last-name"),
		}
	case "person.details":
		f.Context = "The details of a person"
		f.Inputs = []*Field{
			Get("person.name").SetName("name"),
			Get("address").SetName("address"),
		}
	case "name.company", "company.name":
		f.Context = "A full legal name of a company"
		f.Inputs = []*Field{
			Get("string", "160").SetName("registered-name"),
		}
	case "contact.details", "details.contact":
		f.Context = "A full definition of a company"
		f.Inputs = []*Field{
			Get("phone").SetName("phone number"),
			Get("email").SetName("email address"),
			Get("social").SetName("social account"),
		}
	case "company.details", "details.company":
		f.Context = "A full definition of a company"
		f.Inputs = []*Field{
			Get("company.name").SetName("company name"),
			Get("address").SetName("company address"),
		}
	case "phone":
		f.Context = "A phone number"
		f.Element = PHONE
		f.Regexp = `^\+?[1-9]\d{1,14}$` // International phone number format (E.164)
	case "social":
		f.Context = "Social media handle"
		f.Element = STRING
		f.Regexp = `@(\w){1,15}$`
	case "select":
		f.Context = "'dropdown' list"
		f.Element = SELECT
		f.InputOptions = args
		if len(args) == 0 {
			panic("select has no args")
		}
	case "string":
		f.Context = "any string"
		f.Element = STRING
		switch len(args) {
		case 1:
			x, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			f.Range = &Range{}
			f.Range.Min = 1
			f.Range.Max = float64(x)
		case 2:
			x, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(args[1])
			if err != nil {
				panic(err)
			}
			f.Range = &Range{}
			f.Range.Min = float64(x)
			f.Range.Max = float64(y)
		default:
			panic("invalid arg length: " + f.Type)
		}
	default:
		panic("invalid name command")
	}

	f.RegexpHex = hex.EncodeToString([]byte(f.Regexp))

	return
}