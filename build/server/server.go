package main

import (
	"fmt"
	"net/http"
	"github.com/golangdaddy/leap/build/functions"
)

func main() {

	
	http.HandleFunc("/api/project", functions.EntrypointPROJECT)
	http.HandleFunc("/api/projects", functions.EntrypointPROJECTS)
	println("registering handlers for projects")
	
	http.HandleFunc("/api/collection", functions.EntrypointCOLLECTION)
	http.HandleFunc("/api/collections", functions.EntrypointCOLLECTIONS)
	println("registering handlers for collections")
	
	http.HandleFunc("/api/font", functions.EntrypointFONT)
	http.HandleFunc("/api/fonts", functions.EntrypointFONTS)
	println("registering handlers for fonts")
	
	http.HandleFunc("/api/attribute", functions.EntrypointATTRIBUTE)
	http.HandleFunc("/api/attributes", functions.EntrypointATTRIBUTES)
	println("registering handlers for attributes")
	
	http.HandleFunc("/api/layer", functions.EntrypointLAYER)
	http.HandleFunc("/api/layers", functions.EntrypointLAYERS)
	println("registering handlers for layers")
	
	http.HandleFunc("/api/element", functions.EntrypointELEMENT)
	http.HandleFunc("/api/elements", functions.EntrypointELEMENTS)
	println("registering handlers for elements")
	

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}