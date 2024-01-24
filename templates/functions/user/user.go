package main

import (
	"net/http"

	"github.com/golangdaddy/leap/sdk/cloudfunc"
)

// api-user
func (app *App) UserEntrypoint(w http.ResponseWriter, r *http.Request) {

	if cloudfunc.HandleCORS(w, r, "*") {
		return
	}

	function, err := cloudfunc.QueryParam(r, "function")
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	userID, err := cloudfunc.QueryParam(r, "id")
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	user := &User{}
	if err := GetDocument(app.App, userID, user); err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	switch r.Method {

	case "GET":

		switch function {

		case "username":

			if err := cloudfunc.ServeJSON(w, user.Username); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			return

		case "object":

			if err := cloudfunc.ServeJSON(w, user); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			return

		}
	}
}
