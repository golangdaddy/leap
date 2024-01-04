package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/golangdaddy/leap/models"
	"github.com/kr/pretty"
)

type Container struct {
	ProjectID  string
	DatabaseID string
	Object     *models.Object
	Inputs     []string
}

func main() {

	os.RemoveAll("./build/")
	if err := os.Mkdir("./build/", 0777); err != nil {
		panic(err)
	}

	// parse the tree
	folder := "projects/" + os.Args[len(os.Args)-1]
	stack, err := models.ParseStack(folder, "out")
	if err != nil {
		panic(err)
	}
	if stack.DatabaseID == "" {
		panic("set a databaseID")
	}

	if err := copyDir("templates/js/app/", "build/app/"); err != nil {
		panic(err)
	}
	if err := copyDir("templates/js/pages", "build/app/pages/"); err != nil {
		panic(err)
	}
	if err := copyDir("templates/js/dashboard.js", "build/app/features/dashboard.js"); err != nil {
		panic(err)
	}
	if err := copyFile("templates/js/dashboard.js", "build/app/features/dashboard.js"); err != nil {
		panic(err)
	}
	if err := copyFile("templates/js/interfaces.js", "build/app/features/interfaces.js"); err != nil {
		panic(err)
	}
	if err := copyFile("templates/js/home.js", "build/app/features/home.js"); err != nil {
		panic(err)
	}

	// add the entrypoints
	if err := doTemplate("build/app/features/home.js", stack); err != nil {
		panic(err)
	}

	for _, object := range stack.Objects {

		container := Container{
			stack.ProjectID,
			stack.DatabaseID,
			object,
			[]string{},
		}

		required := []string{}
		for _, field := range object.Fields {
			s, err := getInputs(object, field)
			if err != nil {
				println("field name:", field.Name)
				pretty.Println(field)
				panic(err)
			}
			container.Inputs = append(container.Inputs, s)
			if field.Required {
				required = append(required, strings.ToLower(field.Name))
			}
		}
		b, _ := json.Marshal(required)
		container.Inputs = append(
			container.Inputs,
			fmt.Sprintf(`<Submit text="Save" inputs={inputs} submit={props.submit} assert={%s}/>`, string(b)),
		)

		if err := execTemplate(
			"models",
			"model.go",
			strings.ToUpper(object.Name)+".go",
			container,
		); err != nil {
			panic(err)
		}

		if object.HasParent() {
			if err := execTemplate(
				"functions",
				"plural.go",
				strings.ToLower(object.Name)+"s.go",
				container,
			); err != nil {
				panic(err)
			}
		} else {
			if err := execTemplate(
				"functions",
				"pluralNoParent.go",
				strings.ToLower(object.Name)+"s.go",
				container,
			); err != nil {
				panic(err)
			}
		}

		if err := execTemplate(
			"functions",
			"singular.go",
			strings.ToLower(object.Name)+".go",
			container,
		); err != nil {
			panic(err)
		}
		// boilerplater functions
		if err := execTemplate(
			"functions",
			"user.go",
			"user.go",
			container,
		); err != nil {
			panic(err)
		}
		if err := execTemplate(
			"functions",
			"users.go",
			"users.go",
			container,
		); err != nil {
			panic(err)
		}
		if err := execTemplate(
			"functions",
			"auth.go",
			"auth.go",
			container,
		); err != nil {
			panic(err)
		}

		os.MkdirAll(fmt.Sprintf(
			"./build/app/features/%ss",
			cases.Lower(language.English).String(object.Name),
		), 0777)
		os.MkdirAll(fmt.Sprintf(
			"./build/app/features/%ss/forms",
			cases.Lower(language.English).String(object.Name),
		), 0777)
		os.MkdirAll(fmt.Sprintf(
			"./build/app/features/%ss/shared",
			cases.Lower(language.English).String(object.Name),
		), 0777)

		{
			path := fmt.Sprintf(
				"./build/app/features/%ss/%s.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"./templates/js/feature/subject.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				panic(err)
			}
		}
		{
			path := fmt.Sprintf(
				"./build/app/features/%ss/%ss.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"./templates/js/feature/subjects.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				panic(err)
			}
		}
		{
			path := fmt.Sprintf(
				"./build/app/features/%ss/_fetch.js",
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"./templates/js/feature/_fetch.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				panic(err)
			}
		}
		{
			path := fmt.Sprintf(
				"./build/app/features/%ss/_interfaces.js",
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"./templates/js/feature/_interfaces.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				panic(err)
			}
		}
		{
			path := fmt.Sprintf(
				"./build/app/features/%ss/forms/%s.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"./templates/js/feature/forms/subject.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				panic(err)
			}
		}
		{
			path := fmt.Sprintf(
				"./build/app/features/%ss/forms/%sEdit.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"./templates/js/feature/forms/subjectEdit.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				panic(err)
			}
		}
		{
			path := fmt.Sprintf(
				"./build/app/features/%ss/shared/%s.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"./templates/js/feature/shared/subject.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				panic(err)
			}
		}
		{
			path := fmt.Sprintf(
				"./build/app/features/%ss/shared/%sList.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"./templates/js/feature/shared/subjectList.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				panic(err)
			}
		}
		{
			path := fmt.Sprintf(
				"./build/app/features/%ss/new%s.js",
				cases.Lower(language.English).String(object.Name),
				cases.Title(language.English).String(object.Name),
			)
			copyFile(
				"./templates/js/feature/newSubject.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				panic(err)
			}
		}
		{
			path := fmt.Sprintf(
				"./build/app/features/%ss/edit%s.js",
				cases.Lower(language.English).String(object.Name),
				cases.Title(language.English).String(object.Name),
			)
			copyFile(
				"./templates/js/feature/editSubject.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				panic(err)
			}
		}
		{
			path := fmt.Sprintf(
				"./build/app/features/%ss/upload%s.js",
				cases.Lower(language.English).String(object.Name),
				cases.Title(language.English).String(object.Name),
			)
			copyFile(
				"./templates/js/feature/uploadSubject.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				panic(err)
			}
		}
		if object.Mode == "root" {
			stack.Entrypoints = append(stack.Entrypoints, object.Name)
		}
	}

	copyFile(
		"./templates/js/interfaces.js",
		"./build/app/features/interfaces.js",
	)
	if err := doTemplate("./build/app/features/interfaces.js", stack); err != nil {
		panic(err)
	}

	if err := execTemplate(
		"server",
		"server.go",
		"server.go",
		stack,
	); err != nil {
		panic(err)
	}

	copyFile("./templates/models/internals.go", "./build/models/internals.go")
	copyFile("./templates/models/pkg.go", "./build/models/pkg.go")

}

// loaTemplate Parses the template buffer
func loadTemplate(path string) (*template.Template, error) {

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	p := strings.Split(path, "/")
	templateName := p[len(p)-1]

	return template.New(templateName).Funcs(funcMap).Parse(string(b))
}

func execTemplate(folder, fileName, dstName string, data interface{}) error {

	os.MkdirAll("build/"+folder, 0777)

	// Parse the template string
	t, err := loadTemplate(
		fmt.Sprintf("./templates/%s/%s", folder, fileName),
	)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)

	// Execute the template with the provided data
	if err := t.Execute(buf, data); err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("./build/%s/%s", folder, dstName),
		buf.Bytes(),
		0777,
	); err != nil {
		return err
	}

	return nil
}

func doTemplate(path string, data interface{}) error {

	// Parse the template string
	t, err := loadTemplate(path)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)

	// Execute the template with the provided data
	if err := t.Execute(buf, data); err != nil {
		return err
	}

	if err := os.WriteFile(
		path,
		buf.Bytes(),
		0777,
	); err != nil {
		return err
	}

	return nil
}

func copyFile(sourcePath, destinationPath string) error {
	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	fmt.Printf("File %s copied to %s\n", sourcePath, destinationPath)
	return nil
}

func copyDir(src, dest string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		destPath := filepath.Join(dest, path[len(src):])

		if info.IsDir() {
			return os.MkdirAll(destPath, 0777)
		}

		source, err := os.Open(path)
		if err != nil {
			return err
		}
		defer source.Close()

		destination, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		return err
	})
}
