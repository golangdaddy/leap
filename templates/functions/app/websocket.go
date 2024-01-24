package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golangdaddy/leap/sdk/cloudfunc"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Type string
	Body interface{}
}

func (msg *Message) ToJSON() []byte {
	b, _ := json.Marshal(msg)
	return b
}

func (app *App) SendMessageToUser(user *User, msg *Message) {
	app.RLock()
	conn := app.connections[user.Username]
	app.RUnlock()
	if err := conn.WriteMessage(
		websocket.TextMessage,
		msg.ToJSON(),
	); err != nil {
		log.Println(err)
	}
}

func (app *App) HandleConnections(w http.ResponseWriter, r *http.Request) {

	apiKey := r.URL.Query().Get("key")
	if len(apiKey) == 0 {
		err := errors.New("missing key")
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}
	id := app.SeedDigest(apiKey)

	// fetch the Session record
	doc, err := app.Firestore().Collection("sessions").Doc(id).Get(app.Context())
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return
	}
	session := &Session{}
	if err := doc.DataTo(&session); err != nil {
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return
	}

	// fetch the user record
	doc, err = app.Firestore().Collection("users").Doc(session.UserID).Get(app.Context())
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return
	}
	user := &User{}
	if err := doc.DataTo(&user); err != nil {
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return
	}
	defer func() {
		conn.Close()
		app.Lock()
		delete(app.connections, user.Username)
		app.Unlock()
		log.Println("closed connection:", user.Username)
	}()

	app.Lock()
	app.connections[user.Username] = conn
	app.Unlock()

	fmt.Println("Client connected: " + user.Username)

	for range time.NewTicker(time.Minute / 6).C {
		if err := conn.WriteMessage(
			websocket.TextMessage,
			(&Message{
				Type: "shout",
				Body: "hello worlds",
			}).ToJSON(),
		); err != nil {
			log.Println(err)
			return
		}
	}
}
