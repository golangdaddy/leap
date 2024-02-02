package common

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/gin-gonic/gin"
	"github.com/golangdaddy/leap/sdk/assetlayer"
	"github.com/pusher/pusher-http-go/v5"
	"github.com/sashabaranov/go-openai"
)

func (app *App) newClients() Clients {
	return Clients{
		app: app,
	}
}

type Clients struct {
	app        *App
	gin        *gin.Engine
	httpClient *http.Client
	pusher     pusher.Client
	algolia    *search.Client
	assetlayer *assetlayer.Client
	openai     *openai.Client
	sync.RWMutex
}

// UseGin enables a Gin instance
func (self *Clients) Gin() *gin.Engine {

	self.RLock()
	client := self.gin
	self.RUnlock()

	if client == nil {
		self.Lock()
		gin.SetMode(gin.ReleaseMode)
		self.gin = gin.Default()
		defer self.Unlock()
		return self.gin
	}

	return client
}

func (self *Clients) HTTP() *http.Client {

	self.RLock()
	client := self.httpClient
	self.RUnlock()

	if client == nil {
		self.Lock()
		self.httpClient = &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		}
		defer self.Unlock()
		return self.httpClient
	}

	return client
}

func (clients *Clients) Algolia() *search.Client {
	return clients.algolia
}

func (clients *Clients) Assetlayer() *assetlayer.Client {
	return clients.assetlayer
}

func (clients *Clients) ChatGPT() *openai.Client {
	return clients.openai
}

func (clients *Clients) Pusher() pusher.Client {
	return clients.pusher
}
