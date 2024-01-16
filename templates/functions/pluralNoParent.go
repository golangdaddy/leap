{{ $obj := .Object }}
package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/richardboase/npgpublic/sdk/cloudfunc"
	"github.com/richardboase/npgpublic/utils"
	"google.golang.org/api/iterator"
)

// api-{{lowercase .Object.Name}}s
func (app *App) Entrypoint{{uppercase .Object.Name}}S(w http.ResponseWriter, r *http.Request) {

	if cloudfunc.HandleCORS(w, r, "*") {
		return
	}

	_, err := utils.GetSessionUser(app.App, r)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusUnauthorized)
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

		m := map[string]interface{}{}
		if err := cloudfunc.ParseJSON(r, &m); err != nil {
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}

		log.Println("SWITCH FUNCTION", function)

		switch function {

		case "init":

			fields := Fields{{uppercase .Object.Name}}{}
			{{lowercase .Object.Name}} := New{{uppercase .Object.Name}}(nil, fields)
			if !{{lowercase .Object.Name}}.ValidateInput(w, m) {
				return
			}

			// reuse document init create code
			if err := app.CreateDocument{{uppercase .Object.Name}}(nil, {{lowercase .Object.Name}}); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return				
			}

			// finish the request
			if err := cloudfunc.ServeJSON(w, {{lowercase .Object.Name}}); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			return

		{{if eq false .Object.Options.File}}/*{{end}}
		case "initupload":
			// reuse code
			app.Upload{{uppercase .Object.Name}}(w, r, parent)
			return
		{{if eq false .Object.Options.File}}*/{{end}}

		{{if eq false .Object.Options.File}}/*{{end}}
		case "initarchiveupload":
			// reuse code
			app.ArchiveUpload{{uppercase .Object.Name}}(w, r, parent)
			return
		{{if eq false .Object.Options.File}}*/{{end}}


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
				"count": FirestoreCount(app.App, "{{lowercase .Object.Name}}s"),
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

			list := []*{{uppercase .Object.Name}}{}

			// handle objects that need to be ordered
			{{if .Object.Options.Order}}
			q := app.Firestore().Collection("{{lowercase .Object.Name}}s").OrderBy("Meta.Context.Order", firestore.Desc)
			{{else}}
			q := app.Firestore().Collection("{{lowercase .Object.Name}}s").OrderBy("Meta.Modified", firestore.Desc)
			{{end}}
			if limit > 0 {
				q = q.Limit(limit)
			}
			iter := q.Documents(app.Context())
			for {
				doc, err := iter.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					log.Println(err)
					break
				}
				{{lowercase .Object.Name}} := &{{uppercase .Object.Name}}{}
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
