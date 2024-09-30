package leap

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	WebsiteName string
	Config      models.Config
	Object      *models.Object
	Inputs      []string
	EditInputs  []string
}

//go:embed templates/*
//go:embed templates/js/*
//go:embed templates/js/app/*
//go:embed templates/js/feature/*
//go:embed templates/js/pages/*
//go:embed templates/models/*
//go:embed templates/server/*
var f embed.FS

func Build(stack *models.Stack) error {

	if stack.Config.ProjectName == "" {
		panic("set a projectName")
	}

	os.RemoveAll("build/")
	if err := os.MkdirAll("build/app/app", 0777); err != nil {
		return err
	}
	if err := copyDir("templates/js/app", "build/app/"); err != nil {
		return err
	}
	if err := copyDir("templates/js/pages", "build/app/pages/"); err != nil {
		return err
	}
	if err := doTemplate("build/app/pages/login.js", stack); err != nil {
		return err
	}
	if err := doTemplate("build/app/pages/register.js", stack); err != nil {
		return err
	}
	if err := doTemplate("build/app/pages/handcash.js", stack); err != nil {
		return err
	}

	// create models lib and file
	if err := copyFile("templates/models/models.main", "build/models.go"); err != nil {
		return err
	}
	if err := concatModels("build/models.go", stack); err != nil {
		return err
	}
	if err := copyFile("templates/models/models.package", "build/lib/models.go"); err != nil {
		return err
	}
	if err := concatModels("build/lib/models.go", stack); err != nil {
		return err
	}

	println("copying editP")

	if err := copyFile("templates/js/dashboard.js", "build/app/features/dashboard.js"); err != nil {
		return err
	}
	if err := copyFile("templates/js/controller.js", "build/app/features/controller.js"); err != nil {
		return err
	}
	if err := doTemplate("build/app/features/dashboard.js", stack); err != nil {
		return err
	}

	//account
	if err := copyFile("templates/js/account/account.js", "build/app/features/account/account.js"); err != nil {
		return err
	}
	if err := copyFile("templates/js/account/sidebar.js", "build/app/features/account/sidebar.js"); err != nil {
		return err
	}
	if err := copyFile("templates/js/account/accountInbox.js", "build/app/features/account/accountInbox.js"); err != nil {
		return err
	}
	if err := copyFile("templates/js/account/accountInboxCompose.js", "build/app/features/account/accountInboxCompose.js"); err != nil {
		return err
	}
	if err := copyFile("templates/js/account/accountInboxMessages.js", "build/app/features/account/accountInboxMessages.js"); err != nil {
		return err
	}

	if err := copyFile("templates/js/interfaces.js", "build/app/features/interfaces.js"); err != nil {
		return err
	}
	if err := copyFile("templates/js/home.js", "build/app/features/home.js"); err != nil {
		return err
	}

	// assetlayer go
	if err := copyFile("templates/functions/assetlayer/assetlayer.go", "build/api_assetlayer.go"); err != nil {
		return err
	}
	// inbox go
	if err := copyFile("templates/functions/mail/mail.go", "build/api_mail.go"); err != nil {
		return err
	}

	// update the headers and footers
	if err := doTemplate("build/app/components/header.js", stack); err != nil {
		return err
	}
	if err := doTemplate("build/app/components/footer.js", stack); err != nil {
		return err
	}

	// dynamic backend url
	if err := doTemplate("build/app/app/fetch.js", stack); err != nil {
		return err
	}

	objectIndex := map[string]*models.Object{}
	for _, object := range stack.Objects {
		objectIndex[object.Name] = object
	}

	for _, object := range stack.Objects {

		o := *object
		for len(o.Parents) > 0 {
			object.ParentCount++
			o = *objectIndex[o.Parents[0]]
		}

		container := Container{
			stack.WebsiteName,
			stack.Config,
			object,
			[]string{},
			[]string{},
		}

		required := []string{}
		for _, field := range object.GetInputs() {
			if field.Required {
				required = append(required, field.ID)
			}

			s, err := getInputs(object, field)
			if err != nil {
				println("field name:", field.Name)
				pretty.Println(field)
				return err
			}
			container.Inputs = append(container.Inputs, s)

			s, err = getEditInputs(object, field)
			if err != nil {
				println("edit field name:", field.Name)
				pretty.Println(field)
				return err
			}
			container.EditInputs = append(container.EditInputs, s)
		}
		b, _ := json.Marshal(required)
		container.Inputs = append(
			container.Inputs,
			fmt.Sprintf(`<Submit text="Save" inputs={inputs} submit={props.submit} assert={%s}/>`, string(b)),
		)
		container.EditInputs = append(
			container.EditInputs,
			fmt.Sprintf(`<Submit text="Save" inputs={inputs} submit={props.submit} assert={%s}/>`, string(b)),
		)

		// sort handler functions
		if err := execTemplate("functions", "singular.go", "api_"+strings.ToLower(object.Name)+".go", container); err != nil {
			return err
		}
		if err := execTemplate("functions", "singularAdmins.go", "api_"+object.Name+"Admins.go", container); err != nil {
			return err
		}
		if err := execTemplate("functions", "pluralUpload.go", "api_"+object.Name+"sUpload.go", container); err != nil {
			return err
		}
		if err := execTemplate("functions", "pluralLists.go", "api_"+object.Name+"sLists.go", container); err != nil {
			return err
		}
		if err := execTemplate("functions", "pluralCreate.go", "api_"+object.Name+"sCreate.go", container); err != nil {
			return err
		}
		// vertex AI methods
		if err := execTemplate("functions", "pluralShared_VertexCreate.go", "api_"+object.Name+"shared_VertexCreate.go", container); err != nil {
			return err
		}
		// chatgpt methods
		if err := execTemplate("functions", "pluralShared_ChatGPTCreate.go", "api_"+object.Name+"shared_ChatGPTCreate.go", container); err != nil {
			return err
		}
		if err := execTemplate("functions", "pluralShared_ChatGPTEdit.go", "api_"+object.Name+"shared_ChatGPTEdit.go", container); err != nil {
			return err
		}
		if err := execTemplate("functions", "pluralShared_ChatGPTModify.go", "api_"+object.Name+"shared_ChatGPTModify.go", container); err != nil {
			return err
		}
		if err := execTemplate("functions", "pluralShared_ChatGPTPrompt.go", "api_"+object.Name+"shared_ChatGPTPrompt.go", container); err != nil {
			return err
		}
		// plural switch
		if object.HasParent() {
			if err := execTemplate("functions", "plural.go", "api_"+strings.ToLower(object.Name)+"s.go", container); err != nil {
				return err
			}
		} else {
			if err := execTemplate("functions", "pluralNoParent.go", "api_"+strings.ToLower(object.Name)+"s.go", container); err != nil {
				return err
			}
		}

		// terraform templates
		if err := execTemplate("terraform", "terraform.tf", "terraform.tf", container); err != nil {
			return err
		}

		// boilerplater functions
		if err := execTemplate("functions/user", "user.go", "api_user.go", container); err != nil {
			return err
		}
		if err := execTemplate("functions/user", "users.go", "api_users.go", container); err != nil {
			return err
		}
		if err := execTemplate("functions/auth", "utils.go", "utils_auth.go", container); err != nil {
			return err
		}
		if err := execTemplate("functions/auth", "auth.go", "api_auth.go", container); err != nil {
			return err
		}
		if err := execTemplate("functions/auth", "handcash.go", "api_handcash.go", container); err != nil {
			return err
		}
		/*
			os.MkdirAll(fmt.Sprintf(
				"build/app/features/%ss",
				cases.Lower(language.English).String(object.Name),
			), 0777)
			os.MkdirAll(fmt.Sprintf(
				"build/app/features/%ss/forms",
				cases.Lower(language.English).String(object.Name),
			), 0777)
			os.MkdirAll(fmt.Sprintf(
				"build/app/features/%ss/shared",
				cases.Lower(language.English).String(object.Name),
			), 0777)
		*/
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/%s.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/subject.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/%ss.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/subjects.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/_fetch.js",
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/_fetch.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/_interfaces.js",
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/_interfaces.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/forms/%s.js",
				cases.Lower(language.English).String(object.Name),
				"ai",
			)
			copyFile(
				"templates/js/feature/forms/ai.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/forms/%s.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/forms/subject.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/forms/%sEdit.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/forms/subjectEdit.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/assets.js",
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/assets.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/shared/%sAssets.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/shared/subjectAssets.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/shared/%sAssetsRow.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/shared/subjectAssetsRow.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/shared/%sList.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/shared/subjectList.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/shared/%sListRow.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/shared/subjectListRow.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/shared/%sListRowJob.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/shared/subjectListRowJob.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/shared/%sListRowImage.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/shared/subjectListRowImage.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/%ssMatrix.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/subjectsMatrix.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/shared/%sMatrix.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/shared/subjectMatrix.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/shared/%sMatrixRow.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/shared/subjectMatrixRow.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/%sAdmin.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/subjectAdmin.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/%sAdmins.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/subjectAdmins.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/%sMember.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/subjectMember.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/%sMembers.js",
				cases.Lower(language.English).String(object.Name),
				cases.Lower(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/subjectMembers.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/new%s.js",
				cases.Lower(language.English).String(object.Name),
				cases.Title(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/newSubject.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/edit%s.js",
				cases.Lower(language.English).String(object.Name),
				cases.Title(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/editSubject.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/delete%s.js",
				cases.Lower(language.English).String(object.Name),
				cases.Title(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/deleteSubject.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/initUpload%s.js",
				cases.Lower(language.English).String(object.Name),
				cases.Title(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/initUploadSubject.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		{
			path := fmt.Sprintf(
				"build/app/features/%ss/upload%s.js",
				cases.Lower(language.English).String(object.Name),
				cases.Title(language.English).String(object.Name),
			)
			copyFile(
				"templates/js/feature/uploadSubject.js",
				path,
			)
			if err := doTemplate(path, container); err != nil {
				return err
			}
		}
		if len(object.Parents) == 0 {
			stack.Entrypoints = append(stack.Entrypoints, object)
		}

	}

	// add the entrypoints
	if err := doTemplate("build/app/features/home.js", stack); err != nil {
		return err
	}

	copyFile(
		"templates/js/interfaces.js",
		"build/app/features/interfaces.js",
	)
	if err := doTemplate("build/app/features/interfaces.js", stack); err != nil {
		return err
	}

	if err := execTemplate(
		"server",
		"server.go",
		"server.go",
		stack,
	); err != nil {
		return err
	}

	return nil
}

// loaTemplate Parses the template buffer
func loadSourceTemplate(path string) (*template.Template, error) {

	b, err := f.ReadFile(path)
	if err != nil {
		log.Println("loadTemplate failed:" + path)
		return nil, err
	}

	p := strings.Split(path, "/")
	templateName := p[len(p)-1]

	return template.New(templateName).Funcs(funcMap).Parse(string(b))
}

func loadDestinationTemplate(path string) (*template.Template, error) {

	b, err := os.ReadFile(path)
	if err != nil {
		log.Println("loadTemplate failed:" + path)
		return nil, err
	}

	p := strings.Split(path, "/")
	templateName := p[len(p)-1]

	return template.New(templateName).Funcs(funcMap).Parse(string(b))
}

func execTemplate(folder, fileName, dstName string, data interface{}) error {

	os.MkdirAll("build/"+folder, 0777)

	// Parse the template string
	t, err := loadSourceTemplate(
		fmt.Sprintf("templates/%s/%s", folder, fileName),
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
		fmt.Sprintf("build/%s", dstName),
		buf.Bytes(),
		0777,
	); err != nil {
		return err
	}

	return nil
}

func doTemplate(path string, data interface{}) error {

	// Parse the template string
	t, err := loadDestinationTemplate(path)
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

func concatFile(sourcePath, destinationPath string) error {
	fmt.Printf("File %s concat to %s\n", sourcePath, destinationPath)

	s := strings.Split(destinationPath, "/")
	os.MkdirAll(filepath.Join(s[:len(s)-1]...), 0777)
	sourcebytes, err := f.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	destination, err := os.OpenFile(destinationPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = destination.Write(sourcebytes)
	return err
}

func copyFile(sourcePath, destinationPath string) error {
	fmt.Printf("File %s copied to %s\n", sourcePath, destinationPath)

	s := strings.Split(destinationPath, "/")
	os.MkdirAll(filepath.Join(s[:len(s)-1]...), 0777)
	source, err := f.Open(sourcePath)
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

	return nil
}

func copyDir(src, dest string) error {

	println("copying", src, dest)

	//os.MkdirAll(dest, 0777)

	files, err := f.ReadDir(src)
	if err != nil {
		fmt.Printf("error walking the object path %v: %v\n", src, err)
		return err
	}
	for _, file := range files {

		if file.Name() == "node_modules" || file.Name() == ".next" {
			continue
		}

		destPath := filepath.Join(dest, file.Name())

		if file.IsDir() {
			filename := strings.Replace(file.Name(), "/", "", 1)
			if err := copyDir(
				filepath.Join(src, filename),
				filepath.Join(dest, filename),
			); err != nil {
				return err
			}
			continue
		}

		path := filepath.Join(src, file.Name())

		if err := copyFile(path, destPath); err != nil {
			return err
		}
	}

	return nil
}

func concatModels(dstPath string, stack *models.Stack) error {

	if err := concatFile("templates/models/const.go", dstPath); err != nil {
		return err
	}
	if err := concatFile("templates/models/app.go", dstPath); err != nil {
		return err
	}
	if err := concatFile("templates/models/internals.go", dstPath); err != nil {
		return err
	}
	if err := concatFile("templates/models/pusher.go", dstPath); err != nil {
		return err
	}
	if err := concatFile("templates/models/user.go", dstPath); err != nil {
		return err
	}
	if err := concatFile("templates/models/security.go", dstPath); err != nil {
		return err
	}
	if err := concatFile("templates/models/firestore.go", dstPath); err != nil {
		return err
	}
	if err := concatFile("templates/models/assert.go", dstPath); err != nil {
		return err
	}
	if err := concatFile("templates/models/otp.go", dstPath); err != nil {
		return err
	}
	if err := concatFile("templates/models/session.go", dstPath); err != nil {
		return err
	}
	if err := concatFile("templates/models/username.go", dstPath); err != nil {
		return err
	}
	if err := concatFile("templates/models/mail.go", dstPath); err != nil {
		return err
	}
	if err := concatFile("templates/models/asyncjob.go", dstPath); err != nil {
		return err
	}
	if err := concatFile("templates/models/model.go", dstPath); err != nil {
		return err
	}
	if err := doTemplate(dstPath, stack); err != nil {
		return err
	}
	return nil
}
