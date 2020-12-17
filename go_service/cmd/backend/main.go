package main

import (
	"go_service/pkg/api_git"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()

	// r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/", api_git.HomeHandler).Methods("GET")
	r.HandleFunc("/ListFiles", api_git.ListFilesHandler).Methods("POST")
	r.HandleFunc("/CopyRepoFromTag", api_git.CopyRepoFromTagHandler).Methods("POST")
	r.HandleFunc("/CreateMerge", api_git.MergeHandler).Methods("POST")
	r.HandleFunc("/Diff", api_git.DiffHandler).Methods("POST")
	// Start server
	log.Fatal(http.ListenAndServe(":3000", r))

}
