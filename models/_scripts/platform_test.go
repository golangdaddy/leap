package structure

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/golangdaddy/leap"
	"github.com/golangdaddy/leap/models"

	"github.com/otiai10/copy"
)

func TestStructure(t *testing.T) {

	tree := buildStructure(
		models.Config{
			WebsocketHost: "server-go-gen-test-da7z6jf32a-nw.a.run.app",
			WebAPI:        "https://newtown.vercel.app/",
			HostAPI:       "https://server-go-gen-test-da7z6jf32a-nw.a.run.app/",
			RepoURI:       "github.com/golangdaddy/newtown",
			ProjectID:     "npg-generic",
			ProjectName:   "go-gen-test",
			ProjectRegion: "europe-west2-b",
		},
	)

	// Prepare the data model
	if err := models.Prepare(tree); err != nil {
		panic(err)
	}

	// Build the application
	if err := leap.Build(tree); err != nil {
		panic(err)
	}

	// Copy necessary node modules
	if err := copy.Copy("node_modules", "build/app/node_modules"); err != nil {
		log.Println(err)
	}

	// Export debug JSON
	b, err := json.Marshal(tree)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile("../out.json", b, 0755); err != nil {
		panic(err)
	}
}
