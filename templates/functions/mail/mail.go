package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"

	"cloud.google.com/go/firestore"

	"github.com/golangdaddy/leap/sdk/cloudfunc"
	"github.com/kr/pretty"
	"google.golang.org/api/iterator"

	"cloud.google.com/go/language/apiv1beta2/languagepb"
)

// api-inbox
func (app *App) MailEntrypoint(w http.ResponseWriter, r *http.Request) {

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

		// get function
		function, err := cloudfunc.QueryParam(r, "function")
		if err != nil {
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}

		switch function {

		case "convos":

			results := []*Internals{}

			var iter *firestore.DocumentIterator = user.Meta.Firestore(app.App).Collection("convos").Documents(app.Context())
			for {
				doc, err := iter.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					log.Println(err)
					break
				}
				stub := &Internals{}
				if err := doc.DataTo(stub); err != nil {
					log.Println(err)
					continue
				}
				results = append(results, stub)
			}

			if err := cloudfunc.ServeJSON(w, results); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

		case "messages":

			conversation, err := cloudfunc.QueryParam(r, "conversation")
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			results := []*Mail{}

			var iter *firestore.DocumentIterator = app.Firestore().Collection("conversations").Doc(conversation).Collection("messages").Documents(app.Context())
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

		default:
			err := fmt.Errorf("function not found: %s", function)
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return

		}

	case "POST":

		mail := &Mail{}
		if err := cloudfunc.ParseJSON(r, mail); err != nil {
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}
		mail.Meta = user.Meta.NewInternals("mails")
		mail.Sender = user.Ref()
		pretty.Println(mail)

		if len(mail.Body) == 0 {
			log.Println("rejected request with no message body")
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}

		postboxIDs := []string{}
		postboxOwners := []*User{}
		for _, recipient := range append(mail.Recipients, mail.Sender) {
			user, err := app.GetUser(recipient)
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusNotFound)
				return
			}
			postboxOwners = append(postboxOwners, user)
			postboxIDs = append(postboxIDs, user.Meta.ID)
		}

		// make deterministic id from recipients
		sort.Strings(postboxIDs)
		conversation := app.SeedDigest(strings.Join(postboxIDs, " "))

		ownerIDs := []string{}
		for _, owner := range postboxOwners {
			ownerIDs = append(ownerIDs, owner.Username)
		}

		for _, user := range postboxOwners {
			// write the new mail to firestore
			stub := &Internals{
				ID:      conversation,
				Class:   "stubs",
				Updated: true,
				Created: app.TimeNow().Unix(),
				Name:    mail.Body,
			}
			stub.Context.Parents = ownerIDs

			// if this converstation thread between the users exists
			if doc, err := user.Meta.Firestore(app.App).Collection("convos").Doc(conversation).Get(app.Context()); err != nil {
				log.Println(err)
				if _, err := user.Meta.Firestore(app.App).Collection("convos").Doc(conversation).Set(app.Context(), stub); err != nil {
					cloudfunc.HttpError(w, err, http.StatusInternalServerError)
					return
				}
			} else {
				if err := doc.DataTo(stub); err != nil {
					cloudfunc.HttpError(w, err, http.StatusInternalServerError)
					return
				}
				stub.Updated = true
				stub.Modify()
				if _, err := user.Meta.Firestore(app.App).Collection("convos").Doc(conversation).Set(app.Context(), stub); err != nil {
					cloudfunc.HttpError(w, err, http.StatusInternalServerError)
					return
				}
			}

		}

		resp, err := app.GCPClients.NLP().AnalyzeEntitySentiment(app.Context(), &languagepb.AnalyzeEntitySentimentRequest{
			Document: &languagepb.Document{
				Source: &languagepb.Document_Content{
					Content: mail.Body,
				},
				Type: languagepb.Document_PLAIN_TEXT,
			},
		})
		if err != nil {
			cloudfunc.HttpError(w, err, http.StatusInternalServerError)
			return
		}

		pretty.Println(resp)

		// write the new mail to firestore
		if _, err := app.Firestore().Collection("conversations").Doc(conversation).Collection("messages").Doc(mail.Meta.ID).Create(app.Context(), mail); err != nil {
			cloudfunc.HttpError(w, err, http.StatusInternalServerError)
			return
		}

		return

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
