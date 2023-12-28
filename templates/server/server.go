package main

import (
	"fmt"
	"net/http"
	"github.com/golangdaddy/leap/build/functions"
)

func main() {

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