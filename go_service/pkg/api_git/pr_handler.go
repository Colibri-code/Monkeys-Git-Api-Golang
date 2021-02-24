package api_git

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go_service/tools"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PRCreate struct {
	Id               primitive.ObjectID `bson: "_id" `
	NumberPR         int                `bson:"numberPR,omitempty"`         // Consecutive number of the pull request
	IdUser           int                `bson:"idUser,omitempty"`           // id user to create a pull request
	Title            string             `bson:"title,omitempty"`            // Title of thr pull request
	Body             string             `bson:"body,omitempty"`             // Comment of the pull request
	UrlRepoReceivePR string             `bson:"urlRepoReceivePR,omitempty"` // Url to repo the receive the pull request
	UrlRepoCreatePR  string             `bson:"urlRepoCreatePR,omitempty"`  // Url the repo to create a pull request
	CommitHash       string             `bson:"commitHash,omitempty"`       // hash the commit
	Patch            string             `bson:"patch,omitempty"`            // Differences
	BranchNamePR     string             `bson:"branchNamePR,omitempty"`
	IsLocked         bool               `bson:"isLocked,omitempty"`
	Mergeable        bool               `bson:"mergeable,omitempty"`
	HasMerged        bool               `bson:"hasMerged,omitempty"`
	Merged           time.Time          `bson:"merged,omitempty"`
	MergedCommitID   string             `bson:"mergedCommitID,omitempty"`
	MergedBy         int                `bson:"mergedBy,omitempty"`
}

var client *mongo.Client

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

func InsertOne(w http.ResponseWriter, r *http.Request) {

	tools.ConnectionDB()

	var pr PRCreate

	w.Header().Add("content-type", "application/json")

	json.NewDecoder(r.Body).Decode(&pr)

	collection := client.Database("go_git").Collection("PR_Collection")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, _ := collection.InsertOne(ctx, pr)

	json.NewEncoder(w).Encode(result)
}
