package main

import (
	"fmt"
	"log"
	"archive/zip"
	"bytes"
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"encoding/json"

	"github.com/sashabaranov/go-openai"

	{{if eq false .Object.Options.Order}}//{{end}}"google.golang.org/api/iterator"

	"github.com/golangdaddy/leap/sdk/cloudfunc"
)

func (app *App) CreateDocument{{uppercase .Object.Name}}(parent *Internals, object *{{uppercase .Object.Name}}) error {
	log.Println(*object)

	{{if eq false .Object.Options.Order}}/*{{end}}
	var order int
	iter := parent.Firestore(app.App).Collection(object.Meta.Class).Documents(app.Context())
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err)
			break
		}
		order++
	}
	object.Meta.Context.Order = order
	{{if eq false .Object.Options.Order}}*/{{end}}

	{{if eq false .Object.Options.Assetlayer.Wallet}}/*{{end}}
	// create app wallet
	{
		log.Println("CREATING WALLET")
		wallerUserID, err := app.Assetlayer().NewAppWallet(object.Meta.AssetlayerWalletID())
		if err != nil {
			return err
		}
		object.Meta.Wallet = wallerUserID
	}
	{{if eq false .Object.Options.Assetlayer.Wallet}}*/{{end}}

	{{if eq false .Object.Options.Assetlayer.Token}}/*{{end}}
	// create asset
	{
		log.Println("CREATING TOKEN")
		assetID, err := app.Assetlayer().MintAssetWithProperties(object.Meta.AssetlayerCollectionID(), object)
		if err != nil {
			return err
		}
		object.Meta.Asset = assetID
		if err := app.Assetlayer().SendAsset(assetID, "$"+object.Meta.AssetlayerWalletID()); err != nil {
			return err
		}
	}
	{{if eq false .Object.Options.Assetlayer.Token}}*/{{end}}
	
	// write new {{uppercase .Object.Name}} to the DB
	if err := object.Meta.SaveToFirestore(app.App, object); err != nil {
		return err
	}

	{{if eq false .Object.Options.UseCreateTopic}}/*{{end}}
	b, err := app.MarshalJSON(object)
	if err != nil {
		return err
	}
	topicID := "{{.Object.Options.TopicCreate}}"
	result := app.PubSub().Topic(topicID).Publish(
		app.Context(),
		&pubsub.Message{Data: b},
	)
	msgID, err := result.Get(app.Context())
	if err != nil {
		return err
	}
	log.Println("PUBLISHED JOB TO TOPIC", topicID, msgID)
	{{if eq false .Object.Options.UseCreateTopic}}*/{{end}}

	return nil
}

func (app *App) Upload{{uppercase .Object.Name}}(w http.ResponseWriter, r *http.Request, parent *Internals) {

	log.Println("PARSING FORM")
	if err := r.ParseMultipartForm(300 << 20); err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("file")
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	buf := bytes.NewBuffer(nil)
	// Copy the uploaded file to the created file on the filesystem
	if n, err := io.Copy(buf, file); err != nil {
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return
	} else {
		log.Println("copy: wrote", n, "bytes")
	}

	{{if eq false .Object.Options.Image}}/*{{end}}
	if err := checkImage{{uppercase .Object.Name}}(buf.Bytes()); err != nil {
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return
	}
	{{if eq false .Object.Options.Image}}*/{{end}}
	log.Println("creating new {{lowercase .Object.Name}}:", handler.Filename)
	fields := Fields{{uppercase .Object.Name}}{}
	{{lowercase .Object.Name}} := New{{uppercase .Object.Name}}(parent, fields)

	// hidden line here if noparent: {{lowercase .Object.Name}}.Fields.Filename = zipFile.Name
	{{if .Object.Options.File}}{{lowercase .Object.Name}}.Fields.Filename = handler.Filename{{end}}

	// generate a new URI
	uri := {{lowercase .Object.Name}}.Meta.NewURI()
	println ("URI", uri)

	bucketName := "{{.DatabaseID}}-uploads"
	if err := app.write{{titlecase .Object.Name}}File(bucketName, uri, buf.Bytes()); err != nil {
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return
	}

	// reuse document init create code
	if err := app.CreateDocument{{uppercase .Object.Name}}(parent, {{lowercase .Object.Name}}); err != nil {
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return		
	}
	return
}

func (app *App) ArchiveUpload{{uppercase .Object.Name}}(w http.ResponseWriter, r *http.Request, parent *Internals) {

	log.Println("PARSING FORM")
	if err := r.ParseMultipartForm(300 << 20); err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("file")
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	buf := bytes.NewBuffer(nil)
	// Copy the uploaded file to the created file on the filesystem
	if n, err := io.Copy(buf, file); err != nil {
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return
	} else {
		log.Println("copy: wrote", n, "bytes")
	}

	// Open the zip archive from the buffer
	zipReader, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		err = fmt.Errorf("Error opening zip archive: %v", err)
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return 
	}

	// Extract each file from the zip archive
	for n, zipFile := range zipReader.File {

		extractedContent, err := readZipFile{{uppercase .Object.Name}}(zipFile)
		if err != nil {
			cloudfunc.HttpError(w, err, http.StatusInternalServerError)
			return
		}

		{{if eq false .Object.Options.Image}}/*{{end}}
		if err := checkImage{{uppercase .Object.Name}}(extractedContent); err != nil {
			log.Println("skipping file that cannot be decoded:", zipFile.Name)
			continue
		}
		{{if eq false .Object.Options.Image}}*/{{end}}
		log.Println("creating new {{lowercase .Object.Name}}:", zipFile.Name)
		fields := Fields{{uppercase .Object.Name}}{}
		{{lowercase .Object.Name}} := New{{uppercase .Object.Name}}(parent, fields)

		// hidden line here if noparent: {{lowercase .Object.Name}}.Fields.Filename = zipFile.Name
		{{if .Object.Options.File}}{{lowercase .Object.Name}}.Fields.Filename = zipFile.Name{{end}}

		{{lowercase .Object.Name}}.Meta.Context.Order = n

		// generate a new URI
		uri := {{lowercase .Object.Name}}.Meta.NewURI()
		println ("URI", uri)

		bucketName := "{{.DatabaseID}}-uploads"
		if err := app.write{{titlecase .Object.Name}}File(bucketName, uri, extractedContent); err != nil {
			cloudfunc.HttpError(w, err, http.StatusInternalServerError)
			return
		}

		// reuse document init create code
		if err := app.CreateDocument{{uppercase .Object.Name}}(parent, {{lowercase .Object.Name}}); err != nil {
			cloudfunc.HttpError(w, err, http.StatusInternalServerError)
			return		
		}

	}
	return
}

// assert file is an image because of .Object.Options.Image
func checkImage{{uppercase .Object.Name}}(fileBytes []byte) error {
	_, _, err := image.Decode(bytes.NewBuffer(fileBytes))
	return err
}

func readZipFile{{uppercase .Object.Name}}(zipFile *zip.File) ([]byte, error) {
	// Open the file from the zip archive
	zipFileReader, err := zipFile.Open()
	if err != nil {
		return nil, fmt.Errorf("Error opening zip file entry: %v", err)
	}
	defer zipFileReader.Close()

	// Read the content of the file from the zip archive
	var extractedContent bytes.Buffer
	if _, err := io.Copy(&extractedContent, zipFileReader); err != nil {
		return nil, fmt.Errorf("Error reading zip file entry content: %v", err)
	}

	return extractedContent.Bytes(), nil
}

func (app *App) write{{titlecase .Object.Name}}File(bucketName, objectName string, content []byte) error {
	writer := app.GCPClients.GCS().Bucket(bucketName).Object(objectName).NewWriter(app.Context())
	//writer.ObjectAttrs.CacheControl = "no-store"
	defer writer.Close()
	n, err := writer.Write(content)
	fmt.Printf("wrote %s %d bytes to bucket: %s \n", objectName, n, bucketName)
	return err
}

func (app *App) {{lowercase .Object.Name}}ChatGPTCreate(user *User, parent *Internals, prompt string) error {

	fmt.Println("prompt with parent", parent.ID, prompt)

	prompt = fmt.Sprintf(`
ATTENTION! YOUR ENTIRE RESPONSE TO THIS PROMPT NEEDS TO BE A VALID JSON...

We want to create one or more of these data objects: 
{
{{range .Object.Fields}}
	// {{.Context}} {{if .Required}} (THIS FIELD IS REQUIRED){{end}}
	{{lowercase .Name}} ({{lowercase .Type}})
{{end}}
}

The purpose of the object is to represent: {{.Object.Context}}

RULES:
1: USE THIS PROMPT TO GENERATE THE OBJECT OR OBJECT ARRAY: %s
2: GENERATE DATA FOR REQUIRED FIELDS
3: UNLESS SPECIFICALLY TOLD NOT TO, GENERATE ALL FIELDS... DON'T BE LAZY.
4: OMIT ANY NON-REQUIRED FIELDS WHEN NO DATA FOR THE FIELD IS GENERATED.
5: DON'T INCLUDE FIELDS WITH EMPTY STRINGS.
6: RESPECT ANY VALIDATION INFORMATION SPECIFIED FOR FIELDS, SUCH AS MIN AND MAX LENGTHS.
7: REPLY ONLY WITH A JSON ENCODED ARRAY OF THE GENERATED OBJECTS (NO NESTED OBJECTS, JUST OBJECTS IN A JSON ARRAY).
`,
		prompt,
	)

	println(prompt)

	resp, err := app.ChatGPT().CreateChatCompletion(
		app.Context(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		err = fmt.Errorf("ChatCompletion error: %v\n", err)
		return err
	}

	reply := resp.Choices[0].Message.Content
	log.Println("reply >>", reply)

	newResults := []interface{}{}
	if err := json.Unmarshal([]byte(reply), &newResults); err != nil {
		return err
	}

	for _, r := range newResults {
		result, ok := r.(map[string]interface{})
		if !ok {
			return errors.New("array item is not a map")
		}
		// remove any empty fields
		for k, v := range result {
			vv, ok := v.(string)
			if ok && vv == "" {
				delete(result, k)
			}
		}
		object := New{{uppercase .Object.Name}}(parent, Fields{{uppercase .Object.Name}}{})
		if err := object.ValidateObject(result); err != nil {
			return err
		}
		if err := app.CreateDocument{{uppercase .Object.Name}}(parent, object); err != nil {
			return err
		}
		app.SendMessageToUser(user, &Message{Type: "async-create", Body: object})
	}

	return nil
}

func (app *App) {{lowercase .Object.Name}}ChatGPTEdit(user *User, parent *Internals, object *{{uppercase .Object.Name}}, prompt string) error {

	fmt.Println("prompt with parent", parent.ID, prompt)

	objectBytes, err := app.MarshalJSON(object)
	if err != nil {
		return err
	}

	prompt = fmt.Sprintf(`ATTENTION! YOUR ENTIRE RESPONSE TO THIS PROMPT NEEDS TO BE A VALID JSON...

Here is the object we need to edit:
%s

The purpose of the object is to represent: {{.Object.Context}}

RULES:
1: USE THIS PROMPT TO GENERATE THE MUTATION: %s
2: GENERATE DATA FOR REQUIRED FIELDS
3: UNLESS SPECIFICALLY TOLD NOT TO, GENERATE ALL FIELDS... DON'T BE LAZY.
4: OMIT ANY NON-REQUIRED FIELDS WHEN NO DATA FOR THE FIELD IS GENERATED.
5: DON'T INCLUDE FIELDS WITH EMPTY STRINGS.
6: RESPECT ANY VALIDATION INFORMATION SPECIFIED FOR FIELDS, SUCH AS MIN AND MAX LENGTHS.
7: REPLY ONLY WITH THE JSON ENCODED MUTATED OBJECT
`,
		string(objectBytes),
		prompt,
	)

	println(prompt)


	resp, err := app.ChatGPT().CreateChatCompletion(
		app.Context(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		err = fmt.Errorf("ChatCompletion error: %v\n", err)
		return err
	}

	reply := resp.Choices[0].Message.Content
	log.Println("reply >>", reply)

	newResults := []interface{}{}
	if err := json.Unmarshal([]byte(reply), &newResults); err != nil {
		return err
	}

	for _, r := range newResults {
		result, ok := r.(map[string]interface{})
		if !ok {
			return errors.New("array item is not a map")
		}
		// remove any empty fields
		for k, v := range result {
			vv, ok := v.(string)
			if ok && vv == "" {
				delete(result, k)
			}
		}
		object := New{{uppercase .Object.Name}}(parent, Fields{{uppercase .Object.Name}}{})
		if err := object.ValidateObject(result); err != nil {
			return err
		}
		if err := app.CreateDocument{{uppercase .Object.Name}}(parent, object); err != nil {
			return err
		}
		app.SendMessageToUser(user, &Message{Type: "async-create", Body: object})
	}

	return nil
}