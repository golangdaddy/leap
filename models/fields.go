package models

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
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
	ID      string `json:"id"`
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

func (f *Field) Require(name string, args ...string) (g *Field) {
	g = Required(name, args...)
	g.ID = fmt.Sprintf("%s.%s", f.ID, g.ID)
	return
}

func (f *Field) Use(name string, args ...string) (g *Field) {
	g = Get(name, args...)
	g.ID = fmt.Sprintf("%s.%s", f.ID, g.ID)
	return
}

func Required(name string, args ...string) (f *Field) {
	f = Get(name, args...)
	f.Required = true
	return
}

func Get(name string, args ...string) (f *Field) {

	f = &Field{
		Type:    name,
		Element: "none",
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
		f.Element = NUMBER
		f.Range = &Range{}
		f.Range.Min = 0.0
		f.Range.Max = -1.0
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
	case "dimensions":
		f.Context = "the dimensions of the object"
		f.Inputs = []*Field{
			Required("select", "centimeters", "meters", "inches", "feet").SetName("unit").SetCtx("the unit of measurement"),
			Get("float64").SetName("width").SetCtx("the width of an object"),
			Get("float64").SetName("depth").SetCtx("the depth of an object"),
			Get("float64").SetName("height").SetCtx("the height of an object"),
		}
	case "color", "colour":
		f.Element = "colour"
		f.Context = "pick a colour"
	case "phone":
		f.Context = "enter an international phone number"
		f.Element = PHONE
		f.Regexp = `^\+?[1-9]\d{1,14}$` // International phone number format (E.164)
	case "social":
		f.Context = "Social media handle"
		f.Element = STRING
		f.Regexp = `@(\w){1,15}$`
	case "social.account", "account.social":
		f.Context = "Social media account"
		f.Inputs = []*Field{
			Required("select", args...).SetName("social platform").SetCtx("list of social platforms"),
			Required("social").SetName("your handle"),
		}
	case "checkbox":
		f.Context = "a checkbox"
		f.Element = CHECKBOX
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

	f.Name = strings.TrimSpace(strings.ToLower(f.Name))
	f.ID = strings.Replace(f.Name, " ", "-", -1)
	if len(f.ID) == 0 {
		panic("invalid id")
	}
	f.RegexpHex = hex.EncodeToString([]byte(f.Regexp))

	return
}
