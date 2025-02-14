package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

//type requestBody struct {
//	Message string `json:"message"`
//}

var task Task

type JsonStruct struct {
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&task)
	DB.Create(&task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	//fmt.Fprintf(w, "Message is ,%s!", task.Task)
	response := map[string]string{"Задча создана": task.Task}
	json.NewEncoder(w).Encode(response)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	DB.Find(&tasks)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
	//for _, task := range tasks {
	//	fmt.Fprintf(w, "Сообщения из БД %s!\n", task.Task)
	//}
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", PostHandler).Methods("POST")
	router.HandleFunc("/api/messages", HelloHandler).Methods("GET")
	http.ListenAndServe(":8084", router)
}
