package models

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/kr/pretty"
)

//go:embed build.sh
//go:embed _fields/*
//go:embed _objects/*
var f embed.FS

type App struct {
	jsonObjects map[string]*Object
	jsonFields  map[string]*Field
}

func newApp() *App {
	return &App{
		jsonObjects: map[string]*Object{},
		jsonFields:  map[string]*Field{},
	}
}

func (app *App) visitObject(fp string, fi os.DirEntry, err error) error {
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil       // but continue walking elsewhere
	}
	if fi.IsDir() {
		return nil // not a file, ignore.
	}

	name := fi.Name()
	println(name)

	app.jsonObjects[name] = &Object{}

	b, err := f.ReadFile("_objects/" + fp)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, app.jsonObjects[name])
}

func (app *App) visitField(fp string, fi os.DirEntry, err error) error {
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil       // but continue walking elsewhere
	}
	if fi.IsDir() {
		return nil // not a file, ignore.
	}

	name := fi.Name()
	println(name)

	app.jsonFields[name] = &Field{}

	b, err := f.ReadFile("_fields/" + fp)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, app.jsonFields[name])
}

func Prepare(tree *Stack) error {

	app := newApp()

	dir := "test"

	// Walk the directory and call visitFile for each file or directory found
	files, err := f.ReadDir("_objects")
	if err != nil {
		fmt.Printf("error walking the object path %v: %v\n", dir, err)
		return err
	}
	for _, file := range files {
		if err := app.visitObject(file.Name(), file, nil); err != nil {
			return err
		}
	}

	files, err = f.ReadDir("_fields")
	if err != nil {
		fmt.Printf("error walking the field path %v: %v\n", dir, err)
		return err
	}
	for _, file := range files {
		if err := app.visitField(file.Name(), file, nil); err != nil {
			return err
		}
	}

	objectIndex := map[string]*Object{}

	for n, _ := range tree.Objects {

		if tree.Objects[n].Options.File {
			tree.Objects[n].Fields = append(
				tree.Objects[n].Fields,
				&Field{
					Name: "filename",
					JSON: "string_1000",
				},
			)
		}

		if len(tree.Objects[n].JSON) > 0 {
			name := tree.Objects[n].JSON + ".json"
			if app.jsonObjects[name] == nil {
				return errors.New("OBJECT NOT FOUND " + name)
			}
			f := *app.jsonObjects[name]
			f.Name = tree.Objects[n].Name
			f.Parents = tree.Objects[n].Parents
			*tree.Objects[n] = f
		}

		objectIndex[tree.Objects[n].Name] = tree.Objects[n]

		for x, field := range tree.Objects[n].Fields {

			if len(field.JSON) == 0 {
				continue
			}

			name := field.JSON + ".json"
			if app.jsonFields[name] == nil {
				pretty.Println(app.jsonFields)
				return errors.New("FIELD NOT FOUND " + name)
			}
			f := *app.jsonFields[name]
			f.Name = field.Name
			f.Required = field.Required
			*tree.Objects[n].Fields[x] = f

		}
	}

	for _, object := range tree.Objects {
		for _, p := range object.Parents {
			parent := objectIndex[p]
			newObject := *object
			newObject.Fields = nil
			parent.Children = append(parent.Children, &newObject)
		}
	}

	b, err := f.ReadFile("build.sh")
	if err != nil {
		return err
	}
	if err := os.WriteFile("build.sh", b, 0775); err != nil {
		return err
	}

	return nil
}
