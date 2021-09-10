package api_git

import (
	"encoding/json"
	"fmt"
	"go_service/tools"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type MergeRequest struct {
	Url          string `json:"url"`
	Branch       string `json:"branch"`
	TargetBranch string `json:"targetBranch"`
}

func MergeHandler(w http.ResponseWriter, r *http.Request) {
	var mergeRequest MergeRequest
	var response tools.Response
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		response.Message = err.Error()
		response.Result = "Error"
	}
	if body != nil {
		if err := json.Unmarshal(body, &mergeRequest); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

		result, err := CreateMerge(mergeRequest)
		if err != nil {
			response.Message = err.Error()
			response.Result = "Error"
		} else if result != "" {
			response.Message = result
			response.Result = "Exito"
		}

	}

	encodeData, _ := json.Marshal(response)
	fmt.Fprintf(w, string(encodeData))
	return
}

func getBranches(repoUrl string, skip, limit int) ([]string, int, error) {

	repo, err := git.PlainOpen(repoUrl)

	var branchNames []string

	branches, err := repo.Branches()

	if err != nil {
		return nil, 0, err
	}

	i := 0

	count := 0

	_ = branches.ForEach(func(r *plumbing.Reference) error {
		count++

		if i < skip {
			i++
			return nil
		} else if limit != 0 && count > skip+limit {
			return nil
		}

		branchNames = append(branchNames, strings.TrimPrefix(r.Name().String(), BranchPrefix))
		return nil
	})

	return branchNames, count, nil
}
