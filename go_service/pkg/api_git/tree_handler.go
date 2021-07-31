package api_git

import (
	"encoding/json"
	"fmt"
	"go_service/examples"
	"go_service/tools"
	"io"
	"io/ioutil"
	"net/http"
)

type Directory struct {
	Name string `json:"name"`
	File string `json:"file"`
}

func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	var directory Directory
	var response tools.Response
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if body != nil {
		if err := json.Unmarshal(body, &directory); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

		files, err := examples.ListFilesDirectories(directory.Name)
		if err != nil {
			response.Message = err.Error()
			response.Result = "Error"
			encodeData, _ := json.Marshal(response)
			fmt.Fprintf(w, string(encodeData))
			return
		}
		if files != nil {
			encodeData, _ := json.Marshal(files)
			fmt.Fprintf(w, string(encodeData))
			return
		}

	}
}

func ListFileHandler(w http.ResponseWriter, r *http.Request) {
	var directory Directory
	var response tools.Response
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if body != nil {
		if err := json.Unmarshal(body, &directory); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

		files, err := examples.ListFile(directory.Name, directory.File)
		if err != nil {
			response.Message = err.Error()
			response.Result = "Error"
			encodeData, _ := json.Marshal(response)
			fmt.Fprintf(w, string(encodeData))
			return
		}
		if files != "" {
			encodeData, _ := json.Marshal(files)
			fmt.Fprintf(w, string(encodeData))
			return
		}

	}
}

func ListDataToFile(w http.ResponseWriter, r *http.Request) {

	var directory Directory
	var response tools.Response

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if body != nil {
		if err := json.Unmarshal(body, &directory); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}
		files, err := examples.ListBlobFile(directory.Name)
		if err != nil {
			response.Message = err.Error()
			response.Result = "Error"
			encodeData, _ := json.Marshal(response)
			fmt.Fprintf(w, string(encodeData))
			return
		}
		if files != nil {
			encodeData, _ := json.Marshal(files)
			fmt.Fprintf(w, string(encodeData))
			return
		}
	}

}

/*
type Directory struct {
	Name string `json:"name"`
}

func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	var directory Directory
	var response tools.Response
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if body != nil {
		if err := json.Unmarshal(body, &directory); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

		files, err := examples.ListFilesDirectories(directory.Name)
		if err != nil {
			response.Message = err.Error()
			response.Result = "Error"
			encodeData, _ := json.Marshal(response)
			fmt.Fprintf(w, string(encodeData))
			return
		}
		if files != nil {
			encodeData, _ := json.Marshal(files)
			fmt.Fprintf(w, string(encodeData))
			return
		}

	}
}*/
