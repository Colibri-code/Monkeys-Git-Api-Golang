package main

import (
	"encoding/json"
	"fmt"
	"go_service/modules/examples"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/ListFiles", ListFilesHandler).Methods("GET")
	// Start server
	log.Fatal(http.ListenAndServe(":3000", r))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home\n")
}

func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	directory := "./../"
	files, _ := examples.ListFilesDirectories(directory)
	encodeData, _ := json.Marshal(files)
	fmt.Fprintf(w, string(encodeData))
}
