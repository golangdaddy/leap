package main

import (
	"fmt"
	"strconv"
	"net/http"
	"html/template"
	"os"
	"log"
)

func main() {
	log.Println("Starting Application", "{{.ProjectID}}", "{{.ProjectName}}")

	// handle local dev
	if os.Getenv("ENVIRONMENT") != "production" {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/Users/"+os.Getenv("USER")+"/npg-generic-d0985a6033b3.json")
	}

	app := NewApp()

	// init openai
	{{if .Options.ChatGPT}}
	app.UseVertex("{{.ProjectRegion}}")
	app.UseChatGPT(os.Getenv("OPENAI_KEY"))
	{{end}}

	// init pusher
	{{if .Options.Pusher}}
	app.UsePusher(
		os.Getenv("PUSHER_APP_ID"),
		os.Getenv("PUSHER_KEY"),
		os.Getenv("PUSHER_SECRET"),
		os.Getenv("PUSHER_CLUSTER"),
	)
	{{end}}

	// init handcash
	{{if ne nil .Options.Handcash}}
	http.HandleFunc("/handcash/success", app.HandcashEntrypointSuccess)
	{{end}}

	/*
	{{if .Options.Assetlayer}}
	slotID, err := app.Assetlayer().EnsureSlotExists("{{.ProjectName}}-models", "description...", "")
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
	{{end}}
	*/

	http.HandleFunc("/api/user", app.UserEntrypoint)
	http.HandleFunc("/api/users", app.UsersEntrypoint)
	http.HandleFunc("/api/auth", app.AuthEntrypoint)
	http.HandleFunc("/api/mail", app.MailEntrypoint)
	http.HandleFunc("/api/assetlayer", app.EntrypointASSETLAYER)
	{{range .Objects}}
	http.HandleFunc("/api/{{lowercase .Name}}", app.Entrypoint{{uppercase .Name}})
	http.HandleFunc("/api/{{lowercase .Name}}s", app.Entrypoint{{uppercase .Name}}S)
	println("registering handlers for {{lowercase .Name}}s"){{end}}

	http.HandleFunc("/htmx/hello", htmx_hello)
	http.HandleFunc("/htmx/world", htmx_world)

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

func htmx_hello(w http.ResponseWriter, r *http.Request) {
	htmlTemplate := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Hello World</title>
		<script src="https://cdn.jsdelivr.net/npm/htmx.org/dist/htmx.js"></script>
	</head>
	<body>
		<h1 id="hello" hx-put="/htmx/world" hx-target="this" hx-swap="outerHTML"></h1>
		<button hx-get="/world" hx-trigger="click" hx-swap="outerHTML">Refresh</button>
	</body>
	</html>
	`

	// Parse the HTML template
	tmpl, err := template.New("hello").Parse(htmlTemplate)
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}

	// Execute the template and write the output to os.Stdout
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal("Error executing template:", err)
	}
}

func htmx_world(w http.ResponseWriter, r *http.Request) {
	htmlTemplate := `
		<h1 id="hello">hello the world</h1>
`

	// Parse the HTML template
	tmpl, err := template.New("hello").Parse(htmlTemplate)
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}

	// Execute the template and write the output to os.Stdout
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal("Error executing template:", err)
	}
}