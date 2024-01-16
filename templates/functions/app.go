package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/richardboase/npgpublic/sdk/assetlayer"
	"github.com/richardboase/npgpublic/sdk/common"
)

type App struct {
	*common.App
	Assetlayer  *assetlayer.Client
	connections map[string]*websocket.Conn
	sync.RWMutex
}

func NewApp() *App {
	app := &App{
		App: common.NewApp(),
		Assetlayer: assetlayer.NewClient(
			os.Getenv("ASSETLAYERAPPID"),
			os.Getenv("ASSETLAYERSECRET"),
			os.Getenv("DIDTOKEN"),
		),
		connections: map[string]*websocket.Conn{},
	}
	app.UseGCP("{{.ProjectID}}")
	app.UseGCPFirestore("{{.DatabaseID}}")
	return app
}

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

func (app *App) HandleConnections(w http.ResponseWriter, r *http.Request) {

	id := "test"

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		conn.Close()
		app.Lock()
		delete(app.connections, id)
		app.Unlock()
		log.Println("closed connection:", id)
	}()

	app.Lock()
	app.connections[id] = conn
	app.Unlock()

	fmt.Println("Client connected")

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