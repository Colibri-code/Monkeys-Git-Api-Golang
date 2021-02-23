package main

import (
"builtin"
"github.com/gorilla/mux"
"net/http"
"github.com/gorilla/rpc"
"log"
"github.com/Colibri-code/monkeysCloud"

)

routes.HandleFunc(").
Methods("POST")

func CheckArgs(arg ...string) {
	if len(os.Args) < len(arg)+1 {
		Warning("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}


func main (){
	CheckArgs("<ssh-url>", "<directory>", "<tag>", "<name>", "<email>", "<public-key>") 
	url := os.Args[1]
	directory := os.Args[2]
	tag := os.Args[3]
	key := os.Args[6]
	
	r, err := cloneRepo(url, directory, key)

	if err != nil {
		log.Printf("clone repo error: %s", err)
		return
}

    created, err := setTag(r, tag)
	if err != nil {
		log.Printf("create tag error: %s", err)
		return
	}

	if created {
		err = pushTags(r, key)
		if err != nil {
			log.Printf("push tag error: %s", err)
			return
		}
	}
}

type JSONResponse struct {
Fields map[string]string
}

func tagPost(){
    var tagPost bool
    err := r.ParseForm()
    if err != nil {
    log.Println(err.Error)
  }

    res, err := tagPost("github.com/Colibri-code/monkeysCloud")
    if err != nil {
    log.Println(err.Error)
  }

	id, err := res.LastInsertId()
	if err != nil {
	tagPost = false
	} else {
	tagPost = true
	}

createTagPost := strconv.FormatBool(tagPost)
	var resp JSONResponse
	resp.Fields["id"] = string(id)
	resp.Fields["added"] = createTagPost
	jsonResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, jsonResp)
}
