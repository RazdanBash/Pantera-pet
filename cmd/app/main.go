package main

import (
	"github.com/gorilla/mux"
	"myProject/internal/database"
	"myProject/internal/handlers"
	"myProject/internal/taskService"
	"net/http"
)

func main() {
	database.InitDB()

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/api/update/{id}", handler.UpdateTaskHandler).Methods("PATCH")
	http.ListenAndServe(":8084", router)
}
