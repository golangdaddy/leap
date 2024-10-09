package models

import (
	"embed"
	"log"
	"os"
	"strings"

	"github.com/kr/pretty"
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

	pretty.Println(tree.Objects)

	for n := range tree.Objects {

		tree.Objects[n].Name = strings.ToLower(tree.Objects[n].Name)
		if len(tree.Objects[n].Name) == 0 {
			panic("corrupted or missing name")
		}

		log.Println("SETTING INDEX", tree.Objects[n].Name, n)
		objectIndex[tree.Objects[n].Name] = tree.Objects[n]

		// set input field data
		tree.Objects[n].Inputs = tree.Objects[n].GetInputs()

		// normalising data

		// assert or construct plural
		if len(tree.Objects[n].Plural) == 0 {
			tree.Objects[n].Plural = tree.Objects[n].Name + "s"
		}
		tree.Objects[n].Plural = strings.ToLower(tree.Objects[n].Plural)
		if len(tree.Objects[n].Plural) == 0 {
			panic("corrupted or missing plural")
		}

		// make sure all fieldnames are uppercase
		for ii, field := range tree.Objects[n].Fields {
			tree.Objects[n].Fields[ii].Name = strings.ToUpper(field.Name)
		}

		for i := range tree.Objects[n].Names {
			tree.Objects[n].Names[i] = strings.ToUpper(tree.Objects[n].Names[i])
			exists := false
			fields := tree.Objects[n].GetInputs()
			for n, field := range fields {
				field.Name = strings.ToUpper(field.Name)
				println(">>", field.Name)
				if field.Name == tree.Objects[n].Names[i] {
					exists = true
				}
			}
			if !exists {
				panic("can't set reference to field name: " + tree.Objects[n].Names[i])
			}
		}

		for _, p := range tree.Objects[n].Parents {
			parent := objectIndex[p]
			log.Println(p, parent)
			newObject := *tree.Objects[n]
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
