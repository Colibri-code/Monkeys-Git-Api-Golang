package api_git

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type TagRequest struct {
	Url  string `json:"url"`
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
	Result  string `json:"result"`
}

func CopyRepoFromTagHandler(w http.ResponseWriter, r *http.Request) {
	var tagRequest TagRequest
	var response Response
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if body != nil {
		if err := json.Unmarshal(body, &tagRequest); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}
		isCreate, err := CopyRepoFromTag(tagRequest)
		if err != nil && isCreate == false {
			response.Message = err.Error()
			response.Result = "Error"
			encodeData, _ := json.Marshal(response)
			fmt.Fprintf(w, string(encodeData))
			return
		}
		if isCreate {
			response.Message = "Copy success"
			response.Result = "Success"
			encodeData, _ := json.Marshal(response)
			fmt.Fprintf(w, string(encodeData))
			return
		} else {
			response.Message = "Copy failed"
			response.Result = "failed"
			encodeData, _ := json.Marshal(response)
			fmt.Fprintf(w, string(encodeData))
			return
		}
	}
}
