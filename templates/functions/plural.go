{{ $obj := .Object }}
package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"log"

	"github.com/golangdaddy/leap/sdk/cloudfunc"
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

			app.SendMessageToUser(user, "create", object)

			// finish the request
			if err := cloudfunc.ServeJSON(w, object); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			return


		case "upload":
			
			mode, err := cloudfunc.QueryParam(r, "mode")
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			switch mode {
			case "file":
				app.Upload{{uppercase .Object.Name}}(w, r, parent, user)
				return

			case "archive":
				app.ArchiveUpload{{uppercase .Object.Name}}(w, r, parent, user)
				return
			default:
				err := fmt.Errorf("mode not found: %s", mode)
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return		
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
				"count": parent.FirestoreCount(app.App, "{{lowercase .Object.Name}}s"),
			}
			if err := cloudfunc.ServeJSON(w, data); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			return

		case "list":

			mode, err := cloudfunc.QueryParam(r, "mode")
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			var limit int
			limitString, _ := cloudfunc.QueryParam(r, "limit")
			if n, err := strconv.Atoi(limitString); err == nil {
				limit = n
			}

			app.{{lowercase .Object.Name}}Lists(w, user, parent, mode, limit)

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