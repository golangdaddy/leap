package main

import (
	"fmt"
	"strconv"
	"net/http"
	"os"
	"log"
)

func main() {
	log.Println("Starting Application", "{{.ProjectID}}", "{{.DatabaseID}}")

	// handle local dev
	if os.Getenv("ENVIRONMENT") != "production" {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "../../../npg-generic-d0985a6033b3.json")
	}

	app := NewApp()
	app.UseGCP("{{.ProjectID}}")
	app.UseGCPFirestore("{{.DatabaseID}}")
	app.UseAssetlayer(
		os.Getenv("ASSETLAYERAPPID"),
		os.Getenv("ASSETLAYERSECRET"),
		os.Getenv("DIDTOKEN"),
	)
	app.UseChatGPT(os.Getenv("OPENAI_KEY"))

	slotID, err := app.Assetlayer().EnsureSlotExists("{{.DatabaseID}}-models", "description...", "")
	if err != nil {
		panic(err)
	}
	os.Setenv("SLOTID", slotID)

	{{range .Objects}}
	{
		collectionID, err := app.Assetlayer().EnsureCollectionExists(slotID, "Unique", "{{lowercase .Name}}s", "description...", "", 1000000, nil)
		if err != nil {
			panic(err)
		}
		os.Setenv("MODEL_{{uppercase .Name}}S", collectionID)
	}{{end}}

	{{range .Wallets}}
	if _, err := app.Assetlayer().NewAppWallet("{{.}}"); err != nil {
		log.Println(err)
	}{{end}}

	http.HandleFunc("/api/user", app.UserEntrypoint)
	http.HandleFunc("/api/users", app.UsersEntrypoint)
	http.HandleFunc("/api/auth", app.AuthEntrypoint)
	http.HandleFunc("/api/assetlayer", app.EntrypointASSETLAYER)
	http.HandleFunc("/api/asyncjob", app.EntrypointASYNCJOB)
	http.HandleFunc("/api/openai", app.EntrypointOPENAI)
	{{range .Objects}}
	http.HandleFunc("/api/{{lowercase .Name}}", app.Entrypoint{{uppercase .Name}})
	http.HandleFunc("/api/{{lowercase .Name}}s", app.Entrypoint{{uppercase .Name}}S)
	println("registering handlers for {{lowercase .Name}}s"){{end}}

	http.HandleFunc("/ws", app.HandleConnections)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println("Error:", err)
	}
}


