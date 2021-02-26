package api_git

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/ListFiles", ListFilesHandler).Methods("POST")
	r.HandleFunc("/CopyRepoFromTag", CopyRepoFromTagHandler).Methods("POST")
	// Start server
	log.Fatal(http.ListenAndServe(":3000", r))

}
