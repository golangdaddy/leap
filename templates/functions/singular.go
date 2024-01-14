package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/richardboase/npgpublic/sdk/cloudfunc"
	"github.com/richardboase/npgpublic/utils"
)

// api-{{lowercase .Object.Name}}
func (app *App) Entrypoint{{uppercase .Object.Name}}(w http.ResponseWriter, r *http.Request) {

	if cloudfunc.HandleCORS(w, r, "*") {
		return
	}

	_, err := utils.GetSessionUser(app.App, r)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusUnauthorized)
		return
	}

	// get {{lowercase .Object.Name}}
	id, err := cloudfunc.QueryParam(r, "id")
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}
	{{lowercase .Object.Name}} := &{{uppercase .Object.Name}}{}
	if err := utils.GetDocument(app.App, id, {{lowercase .Object.Name}}); err != nil {
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "POST":

		// get function
		function, err := cloudfunc.QueryParam(r, "function")
		if err != nil {
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}

		switch function {

		// update the subject
		case "update":

			m := map[string]interface{}{}
			if err := cloudfunc.ParseJSON(r, &m); err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			if !{{lowercase .Object.Name}}.ValidateInput(w, m) {
				return
			}

			if err := {{lowercase .Object.Name}}.Meta.SaveToFirestore(app.App, {{lowercase .Object.Name}}); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

		case "upload":

			log.Println("PARSING FORM")
			r.ParseMultipartForm(10 << 20)
		
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

			// prepare upload with a new URI
			objectName := {{lowercase .Object.Name}}.Meta.NewURI()
			writer := app.GCPClients.GCS().Bucket("{{.ProjectID}}-uploads").Object(objectName).NewWriter(app.Context())
			//writer.ObjectAttrs.CacheControl = "no-store"
			defer writer.Close()
		
			buf := bytes.NewBuffer(nil)

			// Copy the uploaded file to the created file on the filesystem
			n, err := io.Copy(buf, file)
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			log.Println("UPLOAD copytobuffer: wrote", n, "bytes")

			if _, err := writer.Write(buf.Bytes()); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
		
			// update file with new URI value
			if err := {{lowercase .Object.Name}}.Meta.SaveToFirestore(app.App, {{lowercase .Object.Name}}); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			return
		{{if .Object.Options.Order}}
		case "down":

			list := app.get{{titlecase .Object.Name}}List({{lowercase .Object.Name}})

			var me, beforeMe *{{uppercase .Object.Name}}
			for order, item := range list {
				item.Meta.Context.Order = order
				if item.Meta.ID == {{lowercase .Object.Name}}.Meta.ID {
					me = item
					break
				} else {
					beforeMe = item
				}
			}
			if beforeMe == nil {
				return
			}

			order := me.Meta.Context.Order
			me.Meta.Context.Order = beforeMe.Meta.Context.Order
			beforeMe.Meta.Context.Order = order

			for _, item := range list {
				updates := []firestore.Update{
					{
						FieldPath: firestore.FieldPath{"Meta.Context.Order"},
						Value:     item.Meta.Context.Order,
					},
				}
				println("UPDATING", item.Meta.ID, item.Meta.Context.Order)
				if _, err := item.Meta.Firestore(app.App).Update(app.Context(), updates); err != nil {
					log.Println(err)
				}
			}

			return

		case "up":

			list := app.get{{titlecase .Object.Name}}List({{lowercase .Object.Name}})

			var me, afterMe *{{uppercase .Object.Name}}
			for x, _ := range list {
				order := (len(list) - 1) - x
				item := list[order]
				item.Meta.Context.Order = order
				if item.Meta.ID == {{lowercase .Object.Name}}.Meta.ID {
					me = item
					break
				} else {
					afterMe = item
				}
			}
			if afterMe == nil {
				return
			}

			order := me.Meta.Context.Order
			me.Meta.Context.Order = afterMe.Meta.Context.Order
			afterMe.Meta.Context.Order = order

			for _, item := range list {
				updates := []firestore.Update{
					{
						FieldPath: firestore.FieldPath{"Meta.Context.Order"},
						Value:     item.Meta.Context.Order,
					},
				}
				println("UPDATING", item.Meta.ID, item.Meta.Context.Order)
				if _, err := item.Meta.Firestore(app.App).Update(app.Context(), updates); err != nil {
					log.Println(err)
				}
			}

			return
		{{end}}

		default:
			err := fmt.Errorf("function not found: %s", function)
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}

	case "GET":

		// get function
		function, err := cloudfunc.QueryParam(r, "function")
		if err != nil {
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}

		switch function {

		// return a specific {{lowercase .Object.Name}} object by id
		case "object":

			cloudfunc.ServeJSON(w, {{lowercase .Object.Name}})
			return

		default:
			err := fmt.Errorf("function not found: %s", function)
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}

	case "DELETE":

		_, err := {{lowercase .Object.Name}}.Meta.Firestore(app.App).Delete(app.Context())
		if err != nil {
			cloudfunc.HttpError(w, err, http.StatusInternalServerError)
			return
		}
		return

	default:
		err := errors.New("method not allowed: " + r.Method)
		cloudfunc.HttpError(w, err, http.StatusMethodNotAllowed)
		return
	}
}

func (app *App) get{{titlecase .Object.Name}}List(subject *{{uppercase .Object.Name}}) []*{{uppercase .Object.Name}} {
	list := []*{{uppercase .Object.Name}}{}
	class := subject.Meta.Class
	var ref *firestore.CollectionRef
	if len(subject.Meta.Context.Parent) > 0 {
		ref = Internal(subject.Meta.Context.Parent).Firestore(app.App).Collection(class)
	} else {
		ref = app.Firestore().Collection(class)
	}
	iter := ref.OrderBy("Meta.Context.Order", firestore.Asc).Documents(app.Context())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err)
			break
		}
		object := &{{uppercase .Object.Name}}{}
		if err := doc.DataTo(object); err != nil {
			log.Println(err)
			continue
		}
		list = append(list, object)
	}
	log.Println(len(list))
	return list
}