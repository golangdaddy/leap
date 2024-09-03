package models

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

//go:embed _scripts/*
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

	if err := json.Unmarshal(b, app.jsonObjects[name]); err != nil {
		return err
	}

	println(app.jsonObjects[name].Plural)
	return nil
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

	for n := range tree.Objects {

		if len(tree.Objects[n].JSON) > 0 {
			name := tree.Objects[n].JSON + ".json"
			if app.jsonObjects[name] == nil {
				return errors.New("OBJECT NOT FOUND " + name)
			}

			f := *app.jsonObjects[name]

			if len(f.Name) == 0 {
				f.Name = tree.Objects[n].Name
				if len(f.Name) == 0 {
					panic("f name")
				}
			} else {
				if len(f.Plural) == 0 {
					f.Plural = tree.Objects[n].Plural
				}
				if len(f.Plural) == 0 {
					f.Plural = f.Name + "s"
				}
			}

			f.Parents = tree.Objects[n].Parents
			tree.Objects[n] = &f
		} else {
			if len(tree.Objects[n].Plural) == 0 {
				tree.Objects[n].Plural = tree.Objects[n].Name + "s"
			}
		}

		log.Println("SETTING INDEX", tree.Objects[n].Name, n)
		objectIndex[tree.Objects[n].Name] = tree.Objects[n]

		for x := 0; x < len(tree.Objects[n].Fields); x++ {

			field := tree.Objects[n].Fields[x]

			if len(field.JSON) == 0 {
				continue
			}

			if field.Filter {
				tree.Objects[n].Options.FilterFields = append(tree.Objects[n].Options.FilterFields, field)
			}

			name := field.JSON + ".json"

			if _, ok := app.jsonFields[name]; !ok {
				obj, ok := app.jsonObjects[name]
				if !ok {
					return errors.New("FIELD NOT FOUND " + name)
				}
				println("found object")
				tree.Objects[n].Fields = append(tree.Objects[n].Fields, obj.Fields...)
				continue
			}

			f := *app.jsonFields[name]
			f.Context = field.Context
			f.Name = field.Name
			f.Required = field.Required

			*tree.Objects[n].Fields[x] = f

		}
	}

	for _, object := range tree.Objects {
		for _, p := range object.Parents {
			parent := objectIndex[p]
			log.Println(p, parent)
			newObject := *object
			newObject.Fields = nil
			parent.Children = append(parent.Children, &newObject)
		}
	}

	{
		b, err := f.ReadFile("_scripts/dev.sh")
		if err != nil {
			return err
		}
		if err := os.WriteFile("dev.sh", b, 0775); err != nil {
			return err
		}
	}
	{
		b, err := f.ReadFile("_scripts/Dockerfile")
		if err != nil {
			return err
		}
		if err := os.WriteFile("Dockerfile", b, 0775); err != nil {
			return err
		}
	}

	return nil
}
