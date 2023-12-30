package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/golangdaddy/leap/build/functions"
)

func main() {

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "../../../npg-generic-d0985a6033b3.json")

	http.HandleFunc("/api/user", functions.UserEntrypoint)
	http.HandleFunc("/api/users", functions.UsersEntrypoint)
	http.HandleFunc("/api/auth", functions.AuthEntrypoint)

	{{range .Objects}}
	http.HandleFunc("/api/{{lowercase .Name}}", functions.Entrypoint{{uppercase .Name}})
	http.HandleFunc("/api/{{lowercase .Name}}s", functions.Entrypoint{{uppercase .Name}}S)
	println("registering handlers for {{lowercase .Name}}s")
	{{end}}

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}