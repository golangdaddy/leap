package main

import (
	"fmt"
	"strconv"
	"net/http"
	"os"
	"log"
	"strings"
)

func main() {
	log.Println("Starting Application", "{{.Config.ProjectID}}", "{{.Config.ProjectName}}")

	// handle local dev
	if strings.ToLower(os.Getenv("ENVIRONMENT")) != "production" {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/Users/"+os.Getenv("USER")+"/npg-generic-d0985a6033b3.json")
	}

	app := NewApp()

	// init openai
	{{if .Options.ChatGPT}}
	app.UseVertex("{{.Config.ProjectRegion}}")
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
	{{if .Options.Handcash}}
	http.HandleFunc("/handcash/success", app.HandcashEntrypointSuccess)
	{{end}}


	http.HandleFunc("/api/user", app.UserEntrypoint)
	http.HandleFunc("/api/users", app.UsersEntrypoint)
	http.HandleFunc("/api/auth", app.AuthEntrypoint)
	http.HandleFunc("/api/mail", app.MailEntrypoint)
	http.HandleFunc("/api/assetlayer", app.EntrypointASSETLAYER)
	{{range .Objects}}
	http.HandleFunc("/api/{{lowercase .Name}}", app.Entrypoint{{uppercase .Name}})
	http.HandleFunc("/api/{{.Plural}}", app.Entrypoint{{uppercase .Name}}S)
	println("registering handlers for {{.Plural}}"){{end}}

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
