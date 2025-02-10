package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type requestBody struct {
	Message string `json:"message"`
}

var task string

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody
	json.NewDecoder(r.Body).Decode(&reqBody)
	task = reqBody.Message
	fmt.Fprintf(w, "Есть сообщение : %s", task)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %s", task)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", PostHandler).Methods("POST")
	http.ListenAndServe(":8082", router)
}
