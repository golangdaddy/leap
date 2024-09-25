package models

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type FieldType struct {
	Name  string
	Go    string
	Input string
	Type  string
}

var (
	URL      = &FieldType{"URL", "string", "input", "url"}
	PASSWORD = &FieldType{"PASSWORD", "string", "input", "password"}
	STRING   = &FieldType{"STRING", "string", "input", "text"}
	TEXT     = &FieldType{"TEXT", "string", "textarea", ""}
	INT      = &FieldType{"INT", "int", "input", "number"}
	FLOAT    = &FieldType{"FLOAT", "float64", "input", "number"}
	DATE     = &FieldType{"DATE", "string", "input", "date"}
	TIME     = &FieldType{"TIME", "string", "input", "time"}
	MONTH    = &FieldType{"MONTH", "string", "input", "month"}
	WEEK     = &FieldType{"WEEK", "string", "input", "week"}
	PHONE    = &FieldType{"PHONE", "string", "input", "tel"}
	EMAIL    = &FieldType{"EMAIL", "string", "input", "email"}
	CHECKBOX = &FieldType{"CHECKBOX", "bool", "input", "checkbox"}
	SELECT   = &FieldType{"SELECT", "string", "select", ""}
	COLOR    = &FieldType{"COLOR", "string", "input", "color"}
)

type Field struct {
	ID      string `json:"id"`
	Context string `json:"context"`
	Name    string `json:"name"`
	// go primative types
	Type string `json:"type"`
	// define frontend options
	Element        *FieldType `json:"element"`
	Inputs         []*Field   `json:"inputs,omitempty"`
	InputReference string     `json:"inputReference"`
	InputOptions   []string   `json:"inputOptions,omitempty"`
	Required       bool       `json:"required"`
	Filter         bool       `json:"filter"`
	Range          *Range     `json:"range,omitempty"`
	Regexp         string     `json:"regexp"`
	RegexpHex      string     `json:"regexpHex"`
}

func (object *Object) NewField() *Field {
	return &Field{}
}

type Range struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

func (f *Field) SetName(s string) *Field {
	f.Name, f.ID = cleanName(s)
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
		Type: name,
	}

	switch f.Type {
	case "address":
		f.Context = "An address of a location"
		f.Inputs = []*Field{
			f.Require("int").SetName("building number").SetCtx("the number of the building on the street"),
			f.Use("int").SetName("apartment number").SetCtx("if applicable, the number of the unit or apartment in the building"),
			f.Require("string", "75").SetName("street").SetCtx("the street where the building is"),
			f.Require("string", "50").SetName("town or city").SetCtx("the town or city where the street is"),
			f.Require("string", "50").SetName("country").SetCtx("the country where the town or city is"),
		}
	case "float64":
		f.Context = "64-bit floating-point number"
		f.Element = FLOAT
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
		f.Element = INT
		f.Range = &Range{}
		f.Range.Min = 0.0
		f.Range.Max = -1.0
	case "int":
		f.Context = "any integer"
		f.Element = INT
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
			f.Require("name").SetName("first-name"),
			f.Use("name").SetName("middle-names"),
			f.Require("name").SetName("last-name"),
		}
	case "person.details":
		f.Context = "The details of a person"
		f.Inputs = []*Field{
			f.Use("person.name").SetName("name"),
			f.Use("address").SetName("address"),
		}
	case "name.company", "company.name":
		f.Context = "A full legal name of a company"
		f.Inputs = []*Field{
			f.Use("string", "160").SetName("registered-name"),
		}
	case "contact.details", "details.contact":
		f.Context = "A full definition of a company"
		f.Inputs = []*Field{
			f.Use("phone").SetName("phone number"),
			f.Use("email").SetName("email address"),
			f.Use("social").SetName("social account"),
		}
	case "company.details", "details.company":
		f.Context = "A full definition of a company"
		f.Inputs = []*Field{
			f.Use("company.name").SetName("company name"),
			f.Use("address").SetName("company address"),
		}
	case "dimensions":
		f.Context = "the dimensions of the object"
		f.Inputs = []*Field{
			f.Require("select", "centimeters", "meters", "inches", "feet").SetName("unit").SetCtx("the unit of measurement"),
			f.Use("float64").SetName("width").SetCtx("the width of an object"),
			f.Use("float64").SetName("depth").SetCtx("the depth of an object"),
			f.Use("float64").SetName("height").SetCtx("the height of an object"),
		}
	case "color", "colour":
		f.Element = COLOR
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
			f.Require("select", args...).SetName("social platform").SetCtx("list of social platforms"),
			f.Require("social").SetName("your handle"),
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

	f.Name, f.ID = cleanName(f.Name)

	f.RegexpHex = hex.EncodeToString([]byte(f.Regexp))

	return
}

func cleanName(s string) (string, string) {
	name := strings.TrimSpace(strings.ToLower(s))
	id := strings.Replace(name, " ", "-", -1)
	if len(id) == 0 {
		//panic("invalid id: " + s)
	}
	return name, strings.Replace(strings.ToUpper(id), "-", "", -1)
}
