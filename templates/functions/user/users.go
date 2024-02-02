package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/golangdaddy/leap/sdk/cloudfunc"
	"google.golang.org/api/iterator"
)

// api-users
func (app *App) UsersEntrypoint(w http.ResponseWriter, r *http.Request) {

	if cloudfunc.HandleCORS(w, r, "*") {
		return
	}

	function, err := cloudfunc.QueryParam(r, "function")
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	switch r.Method {

	case "GET":

		switch function {

		case "session":

			user, err := app.GetSessionUser(r)
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusUnauthorized)
				return
			}

			if err := cloudfunc.ServeJSON(w, user); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

		case "autocomplete":

			query, err := cloudfunc.QueryParam(r, "query")
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			// sanitize query
			query = strings.TrimSpace(strings.ToLower(query))

			if len(query) < 3 {
				err := fmt.Errorf("query param is too short: %s", query)
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}
			if len(query) > 14 {
				err := fmt.Errorf("query param is too large: %s", query)
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			list := []UserRef{}

			var iter *firestore.DocumentIterator = app.Firestore().Collection("usernames").Where("Index."+strconv.Itoa(len(query)), "array-contains", query).Limit(20).Documents(context.Background())
			for {
				doc, err := iter.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					log.Println(err)
					break
				}
				username := &Username{}
				if err := doc.DataTo(username); err != nil {
					log.Println(err)
					continue
				}
				list = append(list, username.User)
			}

			if err := cloudfunc.ServeJSON(w, list); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

		}
	}
}
