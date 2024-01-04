package functions

import (
	"context"
	"errors"
	"fmt"
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

		m := map[string]interface{}{}
		if err := cloudfunc.ParseJSON(r, &m); err != nil {
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}

		log.Println("SWITCH FUNCTION", function)

		switch function {

		case "init":

			fields := models.Fields{{uppercase .Object.Name}}{}
			{{lowercase .Object.Name}} := models.New{{uppercase .Object.Name}}(parent, fields)
			if !{{lowercase .Object.Name}}.ValidateInput(w, m) {
				return
			}

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

			q := parent.Firestore(app).Collection("{{lowercase .Object.Name}}s").OrderBy("Meta.Modified", firestore.Desc)
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
