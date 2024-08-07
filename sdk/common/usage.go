package common

import (
	"context"
	"log"
	"os"

	language "cloud.google.com/go/language/apiv1beta2"
	"cloud.google.com/go/vertexai/genai"
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

func (app *App) UseGCPFirestore(projectName string) {
	app.GCPClients.Lock()
	defer app.GCPClients.Unlock()
	app.GCPClients.firestoreDatabase = projectName
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
func (app *App) UseAlgoliaWithVolumeSecret(appID, secretPath string) {
	secretBytes, err := os.ReadFile(secretPath)
	if err != nil {
		log.Fatal(err)
	}
	app.Clients.algolia = search.NewClient(appID, string(secretBytes))
}

// UseAlgolia initialises the algolia client
func (app *App) UseAlgoliaWithEnvSecret(appID, secret string) {
	app.Clients.algolia = search.NewClient(appID, secret)
}

// UsePusher initialises the pusher client
func (app *App) UsePusher(appID, key, secret, cluster string) {
	app.Clients.pusher = &pusher.Client{
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

// google natural language processing
func (self *GCPClients) UseNLP() *language.Client {
	var err error
	self.nlp, err = language.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer self.nlp.Close()
	return self.nlp
}

// google vertex gemini
func (self *GCPClients) UseVertex(location string) *genai.Client {
	self.Lock()
	defer self.Unlock()
	var err error
	self.vertex, err = genai.NewClient(context.Background(), self.projectID, location)
	if err != nil {
		log.Fatal(err)
	}
	return self.vertex
}
