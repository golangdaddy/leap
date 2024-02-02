package main

import (
	"log"
	"net/http"

	"cloud.google.com/go/firestore"

	"github.com/golangdaddy/leap/sdk/cloudfunc"
	"google.golang.org/api/iterator"
)

// api-inbox
func (app *App) Entrypoint(w http.ResponseWriter, r *http.Request) {

	if cloudfunc.HandleCORS(w, r, "*") {
		return
	}

	user, err := app.GetSessionUser(r)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case "GET":

		results := []*Mail{}

		var iter *firestore.DocumentIterator = user.Meta.Firestore(app.App).Collection("inbox").Documents(app.Context())
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Println(err)
				break
			}
			mail := &Mail{}
			if err := doc.DataTo(mail); err != nil {
				log.Println(err)
				continue
			}
			results = append(results, mail)
		}

		if err := cloudfunc.ServeJSON(w, results); err != nil {
			cloudfunc.HttpError(w, err, http.StatusInternalServerError)
			return
		}

	case "POST":

		input := &Mail{}
		if err := cloudfunc.ParseJSON(r, input); err != nil {
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}

		recipient, err := app.GetUser(input.Recipients[0])
		if err != nil {
			cloudfunc.HttpError(w, err, http.StatusNotFound)
			return
		}

		// create mail from user to recipient
		mail := user.NewMail(input.Subject, input.Body, recipient)

		// write the new mail to firestore
		if _, err := app.Firestore().Collection("users").Doc(mail.Recipients[0].ID).Collection("inbox").Doc(mail.Meta.ID).Set(app.Context(), mail); err != nil {
			cloudfunc.HttpError(w, err, http.StatusInternalServerError)
			return
		}

		if err := cloudfunc.ServeJSON(w, mail); err != nil {
			cloudfunc.HttpError(w, err, http.StatusInternalServerError)
			return
		}

	case "DELETE":

		id, err := cloudfunc.QueryParam(r, "id")
		if err != nil {
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}

		if _, err := app.Firestore().Collection("users").Doc(user.Username).Collection("inbox").Doc(id).Delete(app.Context()); err != nil {
			cloudfunc.HttpError(w, err, http.StatusInternalServerError)
			return
		}

	default:

		cloudfunc.HttpError(w, err, http.StatusMethodNotAllowed)

	}
}
