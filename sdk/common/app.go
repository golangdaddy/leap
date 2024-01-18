package common

import (
	"context"
	"crypto/rand"
	"math/big"
	"sync"
	"time"

	"github.com/fxamacker/cbor/v2"
)

type App struct {
	Clients
	AWSClients
	GCPClients
	jwtSigningKey []byte
	cbor          cbor.EncMode
	routes        []Route
	debugMode     bool
	sync.RWMutex
}

func NewApp() *App {

	app := &App{}
	app.newClients()

	return app
}

func (app *App) Context() context.Context {
	return context.Background()
}

func (app *App) Debug(state bool) {
	app.Lock()
	defer app.Unlock()
	app.debugMode = state
}

func (app *App) IsDebug() bool {
	app.RLock()
	defer app.RUnlock()
	return app.debugMode
}

func (app *App) TimeNow() time.Time {
	return time.Now().UTC()
}

func (app *App) RandomNumber(x int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(x)))
	return int(n.Int64())
}
