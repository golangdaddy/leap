{{ $obj := .Object }}
package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/golangdaddy/leap/sdk/cloudfunc"
	"google.golang.org/api/iterator"
)

// api-{{lowercase .Object.Name}}s
func (app *App) Entrypoint{{uppercase .Object.Name}}S(w http.ResponseWriter, r *http.Request) {

	if cloudfunc.HandleCORS(w, r, "*") {
		return
	}

	user, err := app.GetSessionUser(r)
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
	parent, err := app.GetMetadata(parentID)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusNotFound)
		return
	}

	// security
	if !app.IsAdmin(parent, user) {
		err := fmt.Errorf("USER %s IS NOT AN ADMIN", user.Username)
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

		log.Println("SWITCH FUNCTION", function)

		switch function {

		case "openai":

			// get openai command
			mode, err := cloudfunc.QueryParam(r, "mode")
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			m := map[string]interface{}{}
			if err := cloudfunc.ParseJSON(r, &m); err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}
			prompt, ok := AssertSTRING(w, m, "prompt")
			if !ok {
				return
			}

			object := &{{uppercase .Object.Name}}{}
			if err := app.GetDocument(parent.ID, object); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			switch mode {
			case "prompt":

				reply, err := app.{{lowercase .Object.Name}}ChatGPTPrompt(user, object, prompt)
				if err != nil {
					cloudfunc.HttpError(w, err, http.StatusInternalServerError)
					return
				}

				if err := cloudfunc.ServeJSON(w, reply); err != nil {
					cloudfunc.HttpError(w, err, http.StatusInternalServerError)
					return
				}

			case "create":

				if err := app.{{lowercase .Object.Name}}ChatGPTCreate(user, object, prompt); err != nil {
					cloudfunc.HttpError(w, err, http.StatusInternalServerError)
					return
				}

			case "modify":

				if err := app.{{lowercase .Object.Name}}ChatGPTModify(user, object, prompt); err != nil {
					cloudfunc.HttpError(w, err, http.StatusInternalServerError)
					return
				}

			default:
				
				panic("openai")
			}

			return

		case "init":

			m := map[string]interface{}{}
			if err := cloudfunc.ParseJSON(r, &m); err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			fields := Fields{{uppercase .Object.Name}}{}
			object := user.New{{uppercase .Object.Name}}(parent, fields)
			if !object.ValidateInput(w, m) {
				return
			}

			// reuse document init create code
			if err := app.CreateDocument{{uppercase .Object.Name}}(parent, object); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return				
			}

			// finish the request
			if err := cloudfunc.ServeJSON(w, object); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			return

		{{if eq false .Object.Options.File}}/*{{end}}
		case "initupload":
			// reuse code
			app.Upload{{uppercase .Object.Name}}(w, r, parent, user)
			return
		{{if eq false .Object.Options.File}}*/{{end}}

		{{if eq false .Object.Options.File}}/*{{end}}
		case "inituploads":
			// reuse code
			app.ArchiveUpload{{uppercase .Object.Name}}(w, r, parent, user)
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
				"count": parent.FirestoreCount(app.App, "{{lowercase .Object.Name}}s"),
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
			q := parent.Firestore(app.App).Collection("{{lowercase .Object.Name}}s").OrderBy("Meta.Context.Order", firestore.Asc)
			{{else}}
			q := parent.Firestore(app.App).Collection("{{lowercase .Object.Name}}s").OrderBy("Meta.Modified", firestore.Desc)
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