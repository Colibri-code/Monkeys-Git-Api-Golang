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
* Function that insert a pull request to the database
 */
func InsertOne(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")

	var response tools.Response
	//Open Connection with MongoDD
	//Funtion locate in tools Directory
	DBclient := tools.ConnectionDB()

	var Pr PRCreate

	json.NewDecoder(r.Body).Decode(&Pr)

	/*Search the database and collection where the information is stored */

	collection := DBclient.Database("go_git").Collection("PR_collection")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	prCreated, err := pullRequest(Pr)
	if err != nil {
		response.Message = err.Error()
		response.Result = "Error"

	}
	// Insert query from MongoDB
	result, err := collection.InsertOne(ctx, prCreated)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(` {"message":" ` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(result)

	DBclient = tools.Disconnect()

}

/*
* Function that get all from the collection (PR_Collection)
* through the {id} as parameter
 */

func GetAllPr(res http.ResponseWriter, req *http.Request) {

	res.Header().Add("content-type", "application/json")

	//Open Connection with MongoDD
	//Funtion locate in tools Directory
	DBclient := tools.ConnectionDB()

	var PR []PRCreate

	// Stores the name of the database and the collection in the variables (database, PRcollection)

	database := DBclient.Database("go_git")

	PRcollection := database.Collection("PR_collection")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Find Query from MongoDB

	cursor, err := PRcollection.Find(ctx, bson.M{})

	// Error handler
	if err != nil {

		res.Write([]byte(` {"message":" ` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)

	// Scroll through the list to store the information in the PR array
	for cursor.Next(ctx) {
		var PResquest PRCreate
		cursor.Decode(&PResquest)
		PR = append(PR, PResquest)
	}

	//Returns  this line if there are not error
	json.NewEncoder(res).Encode(PR)

	//Disconnect from Database
	DBclient = tools.Disconnect()
}

/*
* Function that get a record from the collection
* through the {id} as parameter
 */
func GetOnePr(res http.ResponseWriter, req *http.Request) {

	res.Header().Add("content-type", "application/json")

	params := mux.Vars(req)

	//Get id params and store in the params variable
	// {Primitive.ObjectIDFromHex} it is because of the type of data that is the id in the database
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var Pr PRCreate

	//Open Connection with MongoDD
	//Funtion locate in tools Directory
	DBclient := tools.ConnectionDB()

	database := DBclient.Database("go_git")

	PRcollection := database.Collection("PR_collection")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// FindOne Query from MongoDB
	//This Query receives 2 parameters
	// The Firts is the context
	// The Second is the Filter (in this case is the {id} variable)

	err := PRcollection.FindOne(ctx, PRCreate{_Id: id}).Decode(&Pr)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(` {"message":" ` + err.Error() + `"}`))
		return
	}

	//Returns  this line if there are not error

	json.NewEncoder(res).Encode(Pr)

	//Disconnect from Database
	DBclient = tools.Disconnect()

}

/*
* Function that update a record from the collection
* through the {id} as parameter
 */
func UpdatePr(res http.ResponseWriter, req *http.Request) {

	res.Header().Add("content-type", "application/json")

	params := mux.Vars(req)

	var Pr PRCreate

	//Get id params and store in the params variable
	// {Primitive.ObjectIDFromHex} it is because of the type of data that is the id in the database
	id, _ := primitive.ObjectIDFromHex(params["id"])

	//Open Connection with MongoDD
	//Funtion locate in tools Directory

	DBclient := tools.ConnectionDB()

	database := DBclient.Database("go_git")

	PRcollection := database.Collection("PR_collection")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	json.NewDecoder(req.Body).Decode(&Pr)

	// UpdateOne Query from MongoDB
	//This Query receives 3 parameters
	// The Firts is the context
	// The Second is the Filter (in this case is the {id} variable)
	// The Third is the data to Update that comes from request(res)
	result, err := PRcollection.UpdateOne(ctx, bson.M{"_id": id}, bson.D{
		{"$set", Pr},
	},
	)
	//Error handler
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(` {"message":" ` + err.Error() + `"}`))
		return
	}

	//Returns  this line if there are not error
	json.NewEncoder(res).Encode(result)

	//Disconnect from Database
	DBclient = tools.Disconnect()
}

/*
* Function that Delete a record from the collection
* through the {id} as parameter
 */

func DeleteOnePr(res http.ResponseWriter, req *http.Request) {

	res.Header().Add("content-type", "application/json")

	params := mux.Vars(req)

	//Get id params and store in the params variable
	// {Primitive.ObjectIDFromHex} it is because of the type of data that is the id in the database

	id, _ := primitive.ObjectIDFromHex(params["id"])

	DBclient := tools.ConnectionDB()

	database := DBclient.Database("go_git")

	PRcollection := database.Collection("PR_collection")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// DeleteOne Query from MongoDB
	//This Query receives 2 parameters
	// The Firts is the context
	// The Second is the Filter (in this case is the {id} variable)
	result, err := PRcollection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(` {"message":" ` + err.Error() + `"}`))
		return
	}
	//Returns  this line if there are not error
	json.NewEncoder(res).Encode(result)

	//Disconnect from Database
	DBclient = tools.Disconnect()
}
