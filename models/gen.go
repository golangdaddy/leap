package models

import (
	"embed"
	"log"
	"os"
)

//go:embed _scripts/*
var f embed.FS

type App struct {
}

func newApp() *App {
	return &App{}
}

func Prepare(tree *Stack) error {

	//app := newApp()

	objectIndex := map[string]*Object{}

	for n := range tree.Objects {

		log.Println("SETTING INDEX", tree.Objects[n].Name, n)
		objectIndex[tree.Objects[n].Name] = tree.Objects[n]

		// set input field data
		tree.Objects[n].Inputs = tree.Objects[n].GetInputs()

		// for x := 0; x < len(tree.Objects[n].Fields); x++ {

		// 	field := tree.Objects[n].Fields[x]

		// 	if len(field.JSON) == 0 {
		// 		continue
		// 	}

		// 	if field.Filter {
		// 		tree.Objects[n].Options.FilterFields = append(tree.Objects[n].Options.FilterFields, field)
		// 	}

		// 	name := field.JSON + ".json"

		// 	if _, ok := app.jsonFields[name]; !ok {
		// 		obj, ok := app.jsonObjects[name]
		// 		if !ok {
		// 			return errors.New("FIELD NOT FOUND " + name)
		// 		}
		// 		println("found object")
		// 		tree.Objects[n].Fields = append(tree.Objects[n].Fields, obj.Fields...)
		// 		continue
		// 	}

		// 	f := *app.jsonFields[name]
		// 	f.Context = field.Context
		// 	f.Name = field.Name
		// 	f.Required = field.Required

		// 	*tree.Objects[n].Fields[x] = f

		// }
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
	if _, err := os.Stat("project.go"); err != nil {
		{
			b, err := f.ReadFile("_scripts/project.go")
			if err != nil {
				return err
			}
			if err := os.WriteFile("project.go", b, 0775); err != nil {
				return err
			}
		}
		{
			b, err := f.ReadFile("_scripts/project_test.go")
			if err != nil {
				return err
			}
			if err := os.WriteFile("project_test.go", b, 0775); err != nil {
				return err
			}
		}
	}
	return nil
}
