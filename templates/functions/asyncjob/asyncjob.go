package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/richardboase/npgpublic/sdk/cloudfunc"
	"github.com/richardboase/npgpublic/utils"
	"google.golang.org/api/iterator"
)

// api-asyncjob
func (app *App) EntrypointASYNCJOB(w http.ResponseWriter, r *http.Request) {

	if cloudfunc.HandleCORS(w, r, "*") {
		return
	}

	_, err := utils.GetSessionUser(app.App, r)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusUnauthorized)
		return
	}

	// get collection metadata
	parentID, err := cloudfunc.QueryParam(r, "parent")
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}
	parent, err := GetMetadata(app.App, parentID)
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

			fields := FieldsASYNCJOB{}
			asyncjob := NewASYNCJOB(parent, fields)

			// reuse document init create code
			if err := app.CreateDocumentASYNCJOB(parent, asyncjob); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			// finish the request
			if err := cloudfunc.ServeJSON(w, asyncjob); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			return

		default:
			err := fmt.Errorf("function not found: %s", function)
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}

	case "GET":

		switch function {

		// return a list of asyncjobs in a specific parent
		case "list":

			var limit int
			limitString, _ := cloudfunc.QueryParam(r, "limit")
			if n, err := strconv.Atoi(limitString); err == nil {
				limit = n
			}

			list := []*ASYNCJOB{}

			// handle objects that need to be ordered

			q := parent.Firestore(app.App).Collection("asyncjobs").OrderBy("Meta.Modified", firestore.Desc)

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
				asyncjob := &ASYNCJOB{}
				if err := doc.DataTo(asyncjob); err != nil {
					log.Println(err)
					continue
				}
				list = append(list, asyncjob)
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

func (app *App) CreateDocumentASYNCJOB(parent *Internals, object *ASYNCJOB) error {
	log.Println(*object)

	// write new ASYNCJOB to the DB
	if err := object.Meta.SaveToFirestore(app.App, object); err != nil {
		return err
	}

	b, err := app.MarshalJSON(object)
	if err != nil {
		return err
	}
	topicID := "{{.DatabaseID}}-jobs"
	result := app.PubSub().Topic(topicID).Publish(
		app.Context(),
		&pubsub.Message{Data: b},
	)
	msgID, err := result.Get(app.Context())
	if err != nil {
		return err
	}
	log.Println("PUBLISHED JOB TO TOPIC", topicID, msgID)

	return nil
}
