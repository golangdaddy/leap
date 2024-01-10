package functions

import (
	"io"
	"context"
	"errors"
	"archive/zip"
	"bytes"
	"fmt"
	{{if .Object.Options.Image}}
	// .Object.Options.Image
	"image"
	_ "image/jpeg"
	_ "image/png"
	{{end}}
	"log"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/richardboase/npgpublic/sdk/cloudfunc"
	"github.com/richardboase/npgpublic/sdk/common"
	"github.com/richardboase/npgpublic/utils"
	"google.golang.org/api/iterator"

	"github.com/golangdaddy/leap/build/models"
)

// api-{{lowercase .Object.Name}}s
func Entrypoint{{uppercase .Object.Name}}S(w http.ResponseWriter, r *http.Request) {

	if cloudfunc.HandleCORS(w, r, "*") {
		return
	}

	ctx := context.Background()

	app := common.NewApp()
	app.UseGCP("{{.ProjectID}}")
	app.UseGCPFirestore("{{.DatabaseID}}")

	_, err := utils.GetSessionUser(app, r)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusUnauthorized)
		return
	}

	{{range $parentID := .Object.Parents}}// get {{$parentID}} metadata{{end}}
	parentID, err := cloudfunc.QueryParam(r, "parent")
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}
	parent, err := models.GetMetadata(app, parentID)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusNotFound)
		return
	}

	// get function
	function, err := cloudfunc.QueryParam(r, "function")
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "POST":

		log.Println("SWITCH FUNCTION", function)

		switch function {

		case "init":

			m := map[string]interface{}{}
			if err := cloudfunc.ParseJSON(r, &m); err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			fields := models.Fields{{uppercase .Object.Name}}{}
			{{lowercase .Object.Name}} := models.New{{uppercase .Object.Name}}(parent, fields)
			if !{{lowercase .Object.Name}}.ValidateInput(w, m) {
				return
			}

			{{if .Object.Options.Order}}
			var order int
			iter := parent.Firestore(app).Collection({{lowercase .Object.Name}}.Meta.Class).Documents(ctx)
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
			{{lowercase .Object.Name}}.Meta.Context.Order = order
			{{end}}

			log.Println(*{{lowercase .Object.Name}})

			// write new {{uppercase .Object.Name}} to the DB
			if err := {{lowercase .Object.Name}}.Meta.SaveToFirestore(app, {{lowercase .Object.Name}}); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			// finish the request
			if err := cloudfunc.ServeJSON(w, {{lowercase .Object.Name}}); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			return

		case "archiveupload":

			log.Println("PARSING FORM")
			if err := r.ParseMultipartForm(300 << 20); err != nil {
				cloudfunc.HttpError(w, err, http.StatusNotFound)
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
				// Open the file from the zip archive
				zipFileReader, err := zipFile.Open()
				if err != nil {
					err = fmt.Errorf("Error opening zip file entry: %v", err)
					cloudfunc.HttpError(w, err, http.StatusInternalServerError)
					return
				}
				defer zipFileReader.Close()
		
				// Read the content of the file from the zip archive
				var extractedContent bytes.Buffer
				if _, err = io.Copy(&extractedContent, zipFileReader); err != nil {
					http.Error(w, fmt.Sprintf("Error reading zip file entry content: %v", err), http.StatusInternalServerError)
					return
				}
		
				// hacky, please investigate
				fileBytes := extractedContent.Bytes()
		
				{{if .Object.Options.Image}}// assert file is an image because of .Object.Options.Image
				_, _, err = image.Decode(bytes.NewBuffer(fileBytes))
				if err != nil {
					log.Println("skipping file that cannot be decoded:", zipFile.Name)
					continue
				}{{end}}
		
				log.Println("creating new {{lowercase .Object.Name}}:", zipFile.Name)
		
				fields := models.Fields{{uppercase .Object.Name}}{}
				{{lowercase .Object.Name}} := models.New{{uppercase .Object.Name}}(parent, fields)

				{{lowercase .Object.Name}}.Meta.Context.Order = n

				// generate a new URI
				uri := {{lowercase .Object.Name}}.Meta.NewURI()
				println ("URI", uri)

				bucketName := "{{.DatabaseID}}-uploads"
				if err := write{{titlecase .Object.Name}}File(app, bucketName, uri, fileBytes); err != nil {
					cloudfunc.HttpError(w, err, http.StatusInternalServerError)
					return
				}
				
				if err := {{lowercase .Object.Name}}.Meta.SaveToFirestore(app, {{lowercase .Object.Name}}); err != nil {
					cloudfunc.HttpError(w, err, http.StatusInternalServerError)
					return
				}

			}
		

		default:
			err := fmt.Errorf("function not found: %s", function)
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}

	case "GET":

		switch function {

		// return the total amount of {{lowercase .Object.Name}}s
		case "count":

			data := map[string]int{
				"count": parent.FirestoreCount(app, "{{lowercase .Object.Name}}s"),
			}
			if err := cloudfunc.ServeJSON(w, data); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			return

		// return a list of {{lowercase .Object.Name}}s in a specific parent
		case "list", "altlist":

			var limit int
			limitString, _ := cloudfunc.QueryParam(r, "limit")
			if n, err := strconv.Atoi(limitString); err == nil {
				limit = n
			}

			list := []*models.{{uppercase .Object.Name}}{}

			// handle objects that need to be ordered
			{{if .Object.Options.Order}}
			q := parent.Firestore(app).Collection("{{lowercase .Object.Name}}s").OrderBy("Meta.Context.Order", firestore.Asc)
			{{else}}
			q := parent.Firestore(app).Collection("{{lowercase .Object.Name}}s").OrderBy("Meta.Modified", firestore.Desc)
			{{end}}

			if limit > 0 {
				q = q.Limit(limit)
			}
			iter := q.Documents(ctx)
			for {
				doc, err := iter.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					log.Println(err)
					break
				}
				{{lowercase .Object.Name}} := &models.{{uppercase .Object.Name}}{}
				if err := doc.DataTo({{lowercase .Object.Name}}); err != nil {
					log.Println(err)
					continue
				}
				list = append(list, {{lowercase .Object.Name}})
			}

			if err := cloudfunc.ServeJSON(w, list); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			return

		default:
			err := fmt.Errorf("function not found: %s", function)
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}

	default:
		err := errors.New("method not allowed: " + r.Method)
		cloudfunc.HttpError(w, err, http.StatusMethodNotAllowed)
		return
	}
}

func write{{titlecase .Object.Name}}File(app *common.App, bucketName, objectName string, content []byte) error {
	writer := app.GCPClients.GCS().Bucket(bucketName).Object(objectName).NewWriter(app.Context())
	//writer.ObjectAttrs.CacheControl = "no-store"
	defer writer.Close()
	n, err := writer.Write(content)
	fmt.Printf("wrote %s %d bytes to bucket: %s \n", objectName, n, bucketName)
	return err
}