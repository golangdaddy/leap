package main

import (
	"sync"

	"github.com/golangdaddy/leap/sdk/common"
	"github.com/gorilla/websocket"
)

type App struct {
	*common.App
	connections map[string]*websocket.Conn
	sync.RWMutex
}

func NewApp() *App {
	app := &App{
		App:         common.NewApp(),
		connections: map[string]*websocket.Conn{},
	}
	return app
}
