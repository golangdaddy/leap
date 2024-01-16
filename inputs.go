package leap

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/golangdaddy/leap/models"
)

func getInputs(object *models.Object, field *models.Field) (string, error) {

	var err error
	var tmp *template.Template

	var output string

	switch field.Input {
	case "select":
		const s = `<Select id="{{lowercase .Name}}" type='text' required={ {{.Required}} } title="%s {{lowercase .Name}}" options={ {{json .InputOptions}} } placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange}/>`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "text":
		const s = `<Input id="{{lowercase .Name}}" type='text' required={ {{.Required}} } title="%s {{lowercase .Name}}" placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange}/>`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "number":
		const s = `<Input id="{{lowercase .Name}}" type='number' required={ {{.Required}} } title="%s {{lowercase .Name}}" inputChange={handleInputChange}/>`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "textarea":
		const s = `<Textarea id="{{lowercase .Name}}" required={ {{.Required}} } title="%s {{lowercase .Name}}" placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange}/>`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "checkbox":
		const s = `<Checkbox id="{{lowercase .Name}}" required={ {{.Required}} } title="%s {{lowercase .Name}}" placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange}/>`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "color":
		const s = `<Color id="{{lowercase .Name}}" required={ {{.Required}} } title="%s {{lowercase .Name}}" inputChange={handleInputChange} />`
		output = fmt.Sprintf(s, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	default:
		return "", fmt.Errorf("missing input for %s %s %s:", object.Name, field.Name, field.Input)
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

	switch field.Input {
	case "select":
		const s = `<Select id="{{lowercase .Name}}" type='text' required={ {{.Required}} } title="%s {{lowercase .Name}}" options={ {{json .InputOptions}} } placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange} value={ inputs["{{lowercase .Name}}"].value } />`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "text":
		const s = `<Input id="{{lowercase .Name}}" type='text' required={ {{.Required}} } title="%s {{lowercase .Name}}" placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange} value={ inputs["{{lowercase .Name}}"].value } />`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "number":
		const s = `<Input id="{{lowercase .Name}}" type='number' required={ {{.Required}} } title="%s {{lowercase .Name}}" inputChange={handleInputChange} value={ inputs["{{lowercase .Name}}"].value } />`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "textarea":
		const s = `<Textarea id="{{lowercase .Name}}" required={ {{.Required}} } title="%s {{lowercase .Name}}" placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange} value={ inputs["{{lowercase .Name}}"].value } />`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "checkbox":
		const s = `<Checkbox id="{{lowercase .Name}}" required={ {{.Required}} } title="%s {{lowercase .Name}}" placeholder="%s {{lowercase .Name}}" inputChange={handleInputChange} value={ inputs["{{lowercase .Name}}"].value } />`
		output = fmt.Sprintf(s, object.Name, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	case "color":
		const s = `<Color id="{{lowercase .Name}}" required={ {{.Required}} } title="%s {{lowercase .Name}}" inputChange={handleInputChange} value={ inputs["{{lowercase .Name}}"].value } />`
		output = fmt.Sprintf(s, object.Name)
		tmp, err = template.New(object.Name + "_" + field.Name).Funcs(funcMap).Parse(output)
	default:
		return "", fmt.Errorf("missing input for %s", field.Input)
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
