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

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PRCreate struct {
	_Id              primitive.ObjectID `json:"_id, omitempty" bson: "_id, omitempty" `
	NumberPR         int                `json:"numberPR, omitempty" bson:"numberPR,omitempty"`                 // Consecutive number of the pull request
	IdUser           int                `json:"idUser, omitempty" bson:"idUser,omitempty"`                     // id user to create a pull request
	Title            string             `json:"title, omitempty" bson:"title,omitempty"`                       // Title of thr pull request
	Body             string             `json:"body, omitempty" bson:"body,omitempty"`                         // Comment of the pull request
	UrlRepoReceivePR string             `json:"urlRepoReceivePR, omitempty" bson:"urlRepoReceivePR,omitempty"` // Url to repo the receive the pull request
	UrlRepoCreatePR  string             `json:"urlRepoCreatePR, omitempty" bson:"urlRepoCreatePR,omitempty"`   // Url the repo to create a pull request
	CommitHash       string             `json:"commitHash, omitempty" bson:"commitHash,omitempty"`             // hash the commit
	Patch            string             `json:"patch, omitempty" bson:"patch,omitempty"`                       // Differences
	BranchNamePR     string             `json:"branchNamePR, omitempty" bson:"branchNamePR,omitempty"`
	IsLocked         bool               `json:"isLocked, omitempty" bson:"isLocked,omitempty"`
	Mergeable        bool               `json:"mergeable, omitempty" bson:"mergeable,omitempty"`
	HasMerged        bool               `json:"hasMerged, omitempty" bson:"hasMerged,omitempty"`
	Merged           time.Time          `json:"merged, omitempty" bson:"merged,omitempty"`
	MergedCommitID   string             `json:"mergedCommitID, omitempty" bson:"mergedCommitID,omitempty"`
	MergedBy         int                `json:"mergedBy, omitempty" bson:"mergedBy,omitempty"`
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

/*
* Function that inserts a pull request to the database
 */
func InsertOne(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")

	DBclient := tools.ConnectionDB()

	var Pr PRCreate

	_ = json.NewDecoder(r.Body).Decode(&Pr)

	/*Look for the Database and PRcollection where its been store*/

	collection := DBclient.Database("go_git").Collection("PR_collection")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := collection.InsertOne(ctx, Pr)
	if err != nil {
		println(err)
	}
	json.NewEncoder(w).Encode(result)

	DBclient = tools.Disconnect()

}

func GetAllPr(res http.ResponseWriter, req *http.Request) {

	res.Header().Add("content-type", "application/json")

	DBclient := tools.ConnectionDB()

	var PR []PRCreate

	database := DBclient.Database("go_git")

	PRcollection := database.Collection("PR_collection")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := PRcollection.Find(ctx, bson.M{})

	if err != nil {

		res.Write([]byte(` {"message":" ` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var PResquest PRCreate
		cursor.Decode(&PResquest)
		PR = append(PR, PResquest)
	}

	json.NewEncoder(res).Encode(PR)

	DBclient = tools.Disconnect()
}

func GetOnePr(res http.ResponseWriter, req *http.Request) {

	res.Header().Add("content-type", "application/json")

	params := mux.Vars(req)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var Pr PRCreate

	DBclient := tools.ConnectionDB()

	database := DBclient.Database("go_git")

	PRcollection := database.Collection("PR_collection")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := PRcollection.FindOne(ctx, PRCreate{_Id: id}).Decode(&Pr)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(` {"message":" ` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(res).Encode(Pr)

	DBclient = tools.Disconnect()

}

func UpdatePr(res http.ResponseWriter, req *http.Request) {

	res.Header().Add("content-type", "application/json")

	params := mux.Vars(req)

	var Pr PRCreate

	id, _ := primitive.ObjectIDFromHex(params["id"])

	DBclient := tools.ConnectionDB()

	database := DBclient.Database("go_git")

	PRcollection := database.Collection("PR_collection")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	json.NewDecoder(req.Body).Decode(&Pr)

	result, err := PRcollection.UpdateOne(ctx, bson.M{"_id": id}, bson.D{
		{"$set", Pr},
	},
	)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(` {"message":" ` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(res).Encode(result)

	DBclient = tools.Disconnect()
}

func DeleteOnePr(res http.ResponseWriter, req *http.Request) {

	res.Header().Add("content-type", "application/json")

	params := mux.Vars(req)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	DBclient := tools.ConnectionDB()

	database := DBclient.Database("go_git")

	PRcollection := database.Collection("PR_collection")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := PRcollection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(` {"message":" ` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(res).Encode(result)

	DBclient = tools.Disconnect()
}
