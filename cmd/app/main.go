package main

import (
	"Pantera-pet/internal/database"
	"github.com/gorilla/mux"
	"internal/handlers"
	"internal/taskService"
	"net/http"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	http.ListenAndServe(":8084", router)
}
