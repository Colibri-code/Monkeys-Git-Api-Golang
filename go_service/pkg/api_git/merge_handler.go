package api_git

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type MergeRequest struct {
	Url          string `json:"url"`
	Branch       string `json:"branch"`
	TargetBranch string `json:"targetBranch"`
}

func MergeHandler(w http.ResponseWriter, r *http.Request) {
	var mergeRequest MergeRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if body != nil {
		if err := json.Unmarshal(body, &mergeRequest); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}
		CreateMerge(mergeRequest)
	}

	fmt.Fprintf(w, "Response\n")

}
