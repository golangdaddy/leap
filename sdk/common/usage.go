package common

import (
	"log"
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/fxamacker/cbor/v2"
	"github.com/golangdaddy/leap/sdk/assetlayer"
	"github.com/pusher/pusher-http-go/v5"
	"github.com/sashabaranov/go-openai"
)

// UseGCP grants the conditions for the GCP services clients
func (app *App) UseGCP(projectID string) {
	if app.debugMode {
		log.Println("CONFIGURING >> GCP with project ", projectID)
	}
	app.GCPClients.Lock()
	defer app.GCPClients.Unlock()
	app.GCPClients.projectID = projectID
	app.GCPClients.firestoreDatabase = "(default)"
}

func (app *App) UseGCPFirestore(databaseID string) {
	app.GCPClients.Lock()
	defer app.GCPClients.Unlock()
	app.GCPClients.firestoreDatabase = databaseID
}

// UseCBOR is an efficient encoding package, check it out
func (app *App) UseCBOR() {
	// setup CBOR encoer
	cb, err := cbor.CanonicalEncOptions().EncMode()
	if err != nil {
		log.Fatalln(err)
	}
	app.cbor = cb
}

// UseAlgolia initialises the algolia client
func (app *App) UseAlgolia(appID, secretPath string) {
	secretBytes, err := os.ReadFile(secretPath)
	if err != nil {
		log.Fatal(err)
	}
	app.Clients.algolia = search.NewClient(appID, string(secretBytes))
}

// UsePusher initialises the pusher client
func (app *App) UsePusher(appID, key, secret, cluster string) {
	app.Clients.pusher = pusher.Client{
		AppID:   appID,
		Key:     key,
		Secret:  secret,
		Cluster: cluster,
	}
}

// UseJWT caches a secret signing key in memory
func (app *App) UseJWT(signingKey string) {
	app.Lock()
	defer app.Unlock()
	app.jwtSigningKey = []byte(signingKey)
}

func (app *App) UseAssetlayer(appID, appSecret, didToken string) {
	app.Lock()
	defer app.Unlock()
	app.Clients.assetlayer = assetlayer.NewClient(
		appID,
		appSecret,
		didToken,
	)
}

// UseGin enables a Gin instance
func (app *App) UseChatGPT(openaiKey string) {
	if len(openaiKey) == 0 {
		panic("OPENAI key not present")
	}
	app.Clients.Lock()
	defer app.Clients.Unlock()
	app.Clients.openai = openai.NewClient(openaiKey)
}
