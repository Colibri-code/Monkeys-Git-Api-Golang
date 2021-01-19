package api_git

import (
	"encoding/json"
	"fmt"
	"go_service/tools"
	"io"
	"io/ioutil"
	"net/http"
)

type DiffRequest struct {
	UrlMain  string `json:"urlMain"`
	HashMain string `json:"hashMain"`
	UrlDiff  string `json:"urlDiff"`
	HashDiff string `json:"hashDiff"`
}

func DiffHandler(w http.ResponseWriter, r *http.Request) {
	var diffRequest DiffRequest
	var response tools.Response
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if body != nil {
		if err := json.Unmarshal(body, &diffRequest); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}
		result, err := diffTreeRepos(diffRequest)
		if err != nil {
			response.Message = err.Error()
			response.Result = "Error"

		}
		if result != nil && err == nil {
			response.Message = result.String()
			response.Result = "Success"
		}

		encodeData, _ := json.Marshal(response)
		fmt.Fprintf(w, string(encodeData))
		return

	} else {
		response.Message = "Body is empty"
		response.Result = "Error"
		encodeData, _ := json.Marshal(response)
		fmt.Fprintf(w, string(encodeData))
	}

}
