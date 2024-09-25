package leap

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/golangdaddy/leap/models"
)

func getInputs(object *models.Object, field *models.Field) (string, error) {

	var err error
	var tmp *template.Template

	var output string

	if field.Element == nil {
		for _, input := range field.Inputs {
			s, err := getInputs(object, input)
			if err != nil {
				return "", err
			}
			output += s
		}
		return output, nil
	}

	switch strings.ToLower(field.Element.Name) {
	case "select":
		const s = `<Select id="{{.ID}}" type='text' required={ {{.Required}} } reference={ "{{.InputReference}}" } referenceParent={ subject } title="%s {{lowercase .Name}}" options={ {{json .InputOptions}} } placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange}/><Spacer/>`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "string", "date", "week", "month", "time", "phone", "email", "color":
		const s = `<Input id="{{.ID}}" type='{{.Type}}' required={ {{.Required}} } title="%s {{lowercase .Name}}" placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange}/><Spacer/>`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "uint":
		const s = `<Input id="{{.ID}}" type='number' required={ {{.Required}} } title="%s {{lowercase .Name}}" inputChange={handleInputChange}/><Spacer/>`
		output = fmt.Sprintf(s, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "int":
		const s = `<Input id="{{.ID}}" type='number' required={ {{.Required}} } title="%s {{lowercase .Name}}" inputChange={handleInputChange}/><Spacer/>`
		output = fmt.Sprintf(s, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "float":
		const s = `<Input id="{{.ID}}" type='number' required={ {{.Required}} } title="%s {{lowercase .Name}}" inputChange={handleInputChange}/><Spacer/>`
		output = fmt.Sprintf(s, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "text":
		const s = `<Textarea id="{{.ID}}" required={ {{.Required}} } title="%s {{lowercase .Name}}" placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange}/><Spacer/>`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "checkbox":
		const s = `<Checkbox id="{{.ID}}" required={ {{.Required}} } title="%s {{lowercase .Name}}" placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange}/><Spacer/>`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "object":
		const s = `<Object id="{{.ID}}" required={ {{.Required}} } title="%s {{lowercase .Name}}" inputChange={handleInputChange} /><Spacer/>`
		output = fmt.Sprintf(s, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "arraystring":
		const s = `<Object id="{{.ID}}" required={ {{.Required}} } title="%s {{lowercase .Name}}" inputChange={handleInputChange} /><Spacer/>`
		output = fmt.Sprintf(s, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	default:
		return "", fmt.Errorf("missing input for %s %s '%s':", object.Name, field.Name, field.Element.Name)
	}
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	// Execute the template with the provided data
	if err := tmp.Execute(buf, field); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func getEditInputs(object *models.Object, field *models.Field) (string, error) {

	var err error
	var tmp *template.Template

	var output string

	if field.Element == nil {
		for _, input := range field.Inputs {
			s, err := getEditInputs(object, input)
			if err != nil {
				return "", err
			}
			output += s
		}
		return output, nil
	}

	switch strings.ToLower(field.Element.Name) {
	case "select":
		const s = `<Select id="{{.ID}}" type='text' required={ {{.Required}} } reference={ "{{.InputReference}}" } referenceParent={ subject } title="%s {{lowercase .Name}}" options={ {{json .InputOptions}} } placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange} value={ inputs["{{.ID}}"].value } /><Spacer/>`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "string", "date", "week", "month", "time", "phone", "email", "color":
		const s = `<Input id="{{.ID}}" type='{{.Type}}' required={ {{.Required}} } title="%s {{lowercase .Name}}" placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange} value={ inputs["{{.ID}}"].value } /><Spacer/>`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "uint":
		const s = `<Input id="{{.ID}}" type='number' required={ {{.Required}} } title="%s {{lowercase .Name}}" inputChange={handleInputChange} value={ inputs["{{.ID}}"].value } /><Spacer/>`
		output = fmt.Sprintf(s, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "int":
		const s = `<Input id="{{.ID}}" type='number' required={ {{.Required}} } title="%s {{lowercase .Name}}" inputChange={handleInputChange} value={ inputs["{{.ID}}"].value } /><Spacer/>`
		output = fmt.Sprintf(s, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "float":
		const s = `<Input id="{{.ID}}" type='number' required={ {{.Required}} } title="%s {{lowercase .Name}}" inputChange={handleInputChange} value={ inputs["{{.ID}}"].value } /><Spacer/>`
		output = fmt.Sprintf(s, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "text":
		const s = `<Textarea id="{{.ID}}" required={ {{.Required}} } title="%s {{lowercase .Name}}" placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange} value={ inputs["{{.ID}}"].value } /><Spacer/>`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "checkbox":
		const s = `<Checkbox id="{{.ID}}" required={ {{.Required}} } title="%s {{lowercase .Name}}" placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange} value={ inputs["{{.ID}}"].value } /><Spacer/>`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "object":
		const s = `<Object id="{{.ID}}" required={ {{.Required}} } title="%s {{lowercase .Name}}" inputChange={handleInputChange} value={ inputs["{{.ID}}"].value } /><Spacer/>`
		output = fmt.Sprintf(s, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "arraystring":
		const s = `<Object id="{{.ID}}" required={ {{.Required}} } title="%s {{lowercase .Name}}" inputChange={handleInputChange} value={ inputs["{{.ID}}"].value } /><Spacer/>`
		output = fmt.Sprintf(s, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	default:
		return "", fmt.Errorf("missing input for %s", field.Element.Name)
	}
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	// Execute the template with the provided data
	if err := tmp.Execute(buf, field); err != nil {
		return "", err
	}

	return buf.String(), nil
}
