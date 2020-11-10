package main

import (
	"log"
	"net/http"

	"./pkg/api_git"

	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()

	// r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/ListFiles", api_git.ListFilesDirectories).Methods("GET")
	// Start server
	log.Fatal(http.ListenAndServe(":3000", r))

}
