package functions

import (
	"io"
	"context"
	"errors"
	"archive/zip"
	"bytes"
	"fmt"
	
	// .Object.Options.Image
	"image"
	_ "image/jpeg"
	_ "image/png"
	
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

// api-elements
func EntrypointELEMENTS(w http.ResponseWriter, r *http.Request) {

	if cloudfunc.HandleCORS(w, r, "*") {
		return
	}

	ctx := context.Background()

	app := common.NewApp()
	app.UseGCP("npg-generic")
	app.UseGCPFirestore("test-project-db")

	_, err := utils.GetSessionUser(app, r)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusUnauthorized)
		return
	}

	// get layer metadata
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

			fields := models.FieldsELEMENT{}
			element := models.NewELEMENT(parent, fields)
			if !element.ValidateInput(w, m) {
				return
			}

			

			log.Println(*element)

			// write new ELEMENT to the DB
			if err := element.Meta.SaveToFirestore(app, element); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			// finish the request
			if err := cloudfunc.ServeJSON(w, element); err != nil {
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
		
				// assert file is an image because of .Object.Options.Image
				_, _, err = image.Decode(bytes.NewBuffer(fileBytes))
				if err != nil {
					log.Println("skipping file that cannot be decoded:", zipFile.Name)
					continue
				}
		
				log.Println("creating new element:", zipFile.Name)
		
				fields := models.FieldsELEMENT{}
				element := models.NewELEMENT(parent, fields)

				element.Meta.Context.Order = n

				// generate a new URI
				uri := element.Meta.NewURI()
				println ("URI", uri)

				bucketName := "test-project-db-uploads"
				if err := writeElementFile(app, bucketName, uri, fileBytes); err != nil {
					cloudfunc.HttpError(w, err, http.StatusInternalServerError)
					return
				}
				
				if err := element.Meta.SaveToFirestore(app, element); err != nil {
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

		// return the total amount of elements
		case "count":

			data := map[string]int{
				"count": parent.FirestoreCount(app, "elements"),
			}
			if err := cloudfunc.ServeJSON(w, data); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			return

		// return a list of elements in a specific parent
		case "list", "altlist":

			var limit int
			limitString, _ := cloudfunc.QueryParam(r, "limit")
			if n, err := strconv.Atoi(limitString); err == nil {
				limit = n
			}

			list := []*models.ELEMENT{}

			// handle objects that need to be ordered
			
			q := parent.Firestore(app).Collection("elements").OrderBy("Meta.Modified", firestore.Desc)
			

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
				element := &models.ELEMENT{}
				if err := doc.DataTo(element); err != nil {
					log.Println(err)
					continue
				}
				list = append(list, element)
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

func writeElementFile(app *common.App, bucketName, objectName string, content []byte) error {
	writer := app.GCPClients.GCS().Bucket(bucketName).Object(objectName).NewWriter(app.Context())
	//writer.ObjectAttrs.CacheControl = "no-store"
	defer writer.Close()
	n, err := writer.Write(content)
	fmt.Printf("wrote %s %d bytes to bucket: %s \n", objectName, n, bucketName)
	return err
}