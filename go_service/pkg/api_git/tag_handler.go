package api_git

import (
	"net/http"
)

type TagRequest struct {
	Url string `json:"url"`
	Tag string `json:"tag"`
}

func CopyRepoFromTagHandler(w http.ResponseWriter, r *http.Request) {

}
