package api_git

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_service/tools"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type PRCreate struct {
	Id               int       `json:"id"`
	NumberPR         int       `json:"numberPR"`         // Consecutive number of the pull request
	IdUser           int       `json:"idUser"`           // id user to create a pull request
	Title            string    `json:"title"`            // Title of thr pull request
	Body             string    `json:"body"`             // Comment of the pull request
	UrlRepoReceivePR string    `json:"urlRepoReceivePR"` // Url to repo the receive the pull request
	UrlRepoCreatePR  string    `json:"urlRepoCreatePR"`  // Url the repo to create a pull request
	CommitHash       string    `json:"commitHash"`       // hash the commit
	Patch            string    `json:"patch"`            // Differences
	BranchNamePR     string    `json:"branchNamePR"`
	IsLocked         bool      `json:"isLocked"`
	Mergeable        bool      `json:"mergeable"`
	HasMerged        bool      `json:"hasMerged"`
	Merged           time.Time `json:"merged"`
	MergedCommitID   string    `json:"mergedCommitID"`
	MergedBy         int       `json:"mergedBy"`
}

func PRHandler(w http.ResponseWriter, r *http.Request) {
	var prCreate PRCreate
	var response tools.Response
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if body != nil {
		if err := json.Unmarshal(body, &prCreate); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
			result, err := pullRequest(prCreate)
			if err != nil {
				response.Message = err.Error()
				response.Result = "Error"

			}
			if result != nil && err == nil {
				sendRequest, _ := json.Marshal(response)
				res, err := http.Post(tools.UrlApi, "application/json", bytes.NewBuffer(sendRequest))
				if res != nil {
					response.Message = "Pull request save"
					response.Result = "Success"
				}
				if err != nil {
					response.Message = err.Error()
					response.Result = "Error"
				}

			}

			encodeData, _ := json.Marshal(response)
			fmt.Fprintf(w, string(encodeData))
			return
		}
	}
}
