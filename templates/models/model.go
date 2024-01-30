{{ $obj := .Object }}
package main

import (
	"log"
	"fmt"
	"errors"
	"net/http"
	"encoding/hex"

	"github.com/golangdaddy/leap/sdk/cloudfunc"
)

func init() {
	// template race fix
	log.Flags()
	hex.DecodeString("FF")
}

type {{uppercase .Object.Name}} struct {
	Meta    Internals
	Fields Fields{{uppercase .Object.Name}} `json:"fields" firestore:"fields"`
}

func New{{uppercase $obj.Name}}(parent *Internals, fields Fields{{uppercase $obj.Name}}) *{{uppercase $obj.Name}} {
	var object *{{uppercase $obj.Name}}
	if parent == nil {
		object = &{{uppercase $obj.Name}}{
			Meta: (Internals{}).NewInternals("{{lowercase $obj.Name}}s"),
			Fields: fields,
		}
	} else {
		object = &{{uppercase $obj.Name}}{
			Meta: parent.NewInternals("{{lowercase $obj.Name}}s"),
			Fields: fields,
		}
	}
	object.Meta.Context.Children = []string{
		{{range .Object.Children}}"{{.Name}}",
		{{end}}
	}
	return object
}

type Fields{{uppercase .Object.Name}} struct {
	{{range .Object.Fields}}{{titlecase .Name}} {{.Type}} `json:"{{lowercase .Name}}" firestore:"{{lowercase .Name}}"`
	{{end}}
}

func (x *{{uppercase .Object.Name}}) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {
	if err := x.ValidateObject(m); err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return false
	}
	return true
}

func (x *{{uppercase .Object.Name}}) ValidateObject(m map[string]interface{}) error {

	var err error
	var exists bool
	{{range .Object.Fields}}

	_, exists = m["{{lowercase .Name}}"]
	if {{.Required}} && !exists {
		return errors.New("required field '{{lowercase .Name}}' not supplied")
	}
	if exists {
		x.Fields.{{titlecase .Name}}, err = assert{{uppercase .Type}}(m, "{{lowercase .Name}}")
		if err != nil {
			return errors.New(err.Error())
		}
		{
			exp := "{{.Regexp}}"
			if len(exp) > 0 {
				if !RegExp(exp, fmt.Sprintf("%v", x.Fields.{{titlecase .Name}})) {
					return fmt.Errorf("failed to regexp: %s >> %s", exp, x.Fields.{{titlecase .Name}})
				}
			}
		}
		{
			exp := "{{.RegexpHex}}"
			if len(exp) > 0 {
				log.Println("EXPR", exp)
				b, err := hex.DecodeString(exp)
				if err != nil {
					log.Println(err)
				}
				if !RegExp(string(b), fmt.Sprintf("%v", x.Fields.{{titlecase .Name}})) {
					return fmt.Errorf("failed to regexpHex: %s >> %s", string(b), x.Fields.{{titlecase .Name}})
				}
			}
		}
		{{if .Range}}
		if err := assertRangeMin({{.Range.Min}}, x.Fields.{{titlecase .Name}}); err != nil {
			{{if .Required}}
			return err
			{{end}}
		}
		if err := assertRangeMax({{.Range.Max}}, x.Fields.{{titlecase .Name}}); err != nil {
			return err
		}
		{{end}}
	}
	{{end}}

	x.Meta.Modify()

	return nil
}
/*
func (x *{{uppercase .Object.Name}}) ValidateByCount(w http.ResponseWriter, m map[string]interface{}, count int) bool {

	var counter int
	var exists bool
	{{range .Object.Fields}}
	x.Fields.{{titlecase .Name}}, exists = Assert{{uppercase .Type}}(w, m, "{{lowercase .Name}}")
	if exists {
		counter++
	}

	{
		exp := "{{.Regexp}}"
		if len(exp) > 0 {
			if !RegExp(exp, fmt.Sprintf("%v", x.Fields.{{titlecase .Name}})) {
				return fmt.Errorf("failed to regexp: %s >> %s", exp, x.Fields.{{titlecase .Name}})
			}
		}
	}
	{
		exp := "{{.RegexpHex}}"
		if len(exp) > 0 {
			log.Println("EXPR", exp)
			b, err := hex.DecodeString(exp)
			if err != nil {
				log.Println(err)
			}
			if !RegExp(string(b), fmt.Sprintf("%v", x.Fields.{{titlecase .Name}})) {
				return fmt.Errorf("failed to regexpHex: %s >> %s", string(b), x.Fields.{{titlecase .Name}})
			}
		}
	}

	{{if .Range}}
	{{if .Required}}
	if !AssertRangeMin(w, {{.Range.Min}}, x.Fields.{{titlecase .Name}}) {
		return false
	}
	{{end}}
	if !AssertRangeMax(w, {{.Range.Max}}, x.Fields.{{titlecase .Name}}) {
		return false
	}
	{{end}}{{end}}

	x.Meta.Modify()

	return counter == count
}
*/