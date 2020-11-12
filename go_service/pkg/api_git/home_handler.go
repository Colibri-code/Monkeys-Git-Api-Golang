package api_git

import (
	"encoding/json"
	"fmt"
	"go_service/examples"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	directory := "./../"
	files, _ := examples.ListFilesDirectories(directory)
	encodeData, _ := json.Marshal(files)
	fmt.Fprintf(w, string(encodeData))
}
