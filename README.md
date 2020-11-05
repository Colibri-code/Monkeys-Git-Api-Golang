# Instructions

Install Docker https://www.docker.com/products/docker-desktop

Install Docker compose

Run docker-compose build

Run docker-compose up

In visual studio code attach to running container.



### Gorilla Mux 
You can see all the documentation of the framework in https://github.com/gorilla/mux#examples 
 By default, the request is a GET when you create a Handle function and attach it to a route, if you want to specify you have to add at the end of the r.HandleFunc("/",yourhandler) the function Methods() and passing the type or types of request you want “GET” “POST” “PUT” or "DELETE"
Example:
```go
r.HandleFunc("/articles/{id}", yourhandler).Methods("GET")
```
Your handleFunction
```go
func yourHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "ID: %v\n", vars["id"])
}
```


