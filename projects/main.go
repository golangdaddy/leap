package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golangdaddy/leap/models"
)

type App struct {
	jsonObjects map[string]*models.Object
	jsonFields  map[string]*models.Field
}

func newApp() *App {
	return &App{
		jsonObjects: map[string]*models.Object{},
		jsonFields:  map[string]*models.Field{},
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

	app.jsonObjects[name] = &models.Object{}

	b, err := os.ReadFile(fp)
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

	app.jsonFields[name] = &models.Field{}

	b, err := os.ReadFile(fp)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, app.jsonFields[name])
}

func main() {

	app := newApp()

	dir := "test"

	// Walk the directory and call visitFile for each file or directory found
	if err := filepath.WalkDir("_objects", app.visitObject); err != nil {
		fmt.Printf("error walking the object path %v: %v\n", dir, err)
	}
	if err := filepath.WalkDir("_fields", app.visitField); err != nil {
		fmt.Printf("error walking the field path %v: %v\n", dir, err)
	}

	folder := os.Args[len(os.Args)-1]
	tree, err := models.ParseStack(folder, "tree")
	if err != nil {
		panic(err)
	}

	objectIndex := map[string]*models.Object{}

	for n, _ := range tree.Objects {

		if len(tree.Objects[n].JSON) > 0 {
			name := tree.Objects[n].JSON + ".json"
			if app.jsonObjects[name] == nil {
				panic("OBJECT NOT FOUND " + name)
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
				panic("FIELD NOT FOUND " + name)
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

	b, err := json.Marshal(tree)
	if err != nil {
		panic(err)
	}

	os.WriteFile(folder+"/out.json", b, 0775)
}
