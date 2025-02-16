package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//var task Task

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	DB.Create(&task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	task := Task{}
	vars := mux.Vars(r)
	key := vars["id"]
	id, err := strconv.Atoi(key)
	{
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
	//deleted := task
	json.NewDecoder(r.Body).Decode(&task)
	DB.Delete(&Task{}, id)
	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Content-Type", "application/json")
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	vars := mux.Vars(r)
	id := vars["id"]
	DB.First(&task, id)

	var updateData struct {
		Task   string `json:"task"`
		IsDone bool   `json:"is_done"`
	}
	json.NewDecoder(r.Body).Decode(&updateData)
	task.Task = updateData.Task
	task.IsDone = updateData.IsDone

	DB.Model(&task).Updates(map[string]interface{}{
		"task":    updateData.Task,
		"is_done": updateData.IsDone,
	})

	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(map[string]interface{}{"id": task.ID, "updated": true})
	json.NewEncoder(w).Encode(task)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var task []Task
	DB.Order("id asc").Find(&task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

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
	router.HandleFunc("/api/delete/{id}", DeleteHandler).Methods("DELETE")
	router.HandleFunc("/api/patch/{id}", PatchHandler).Methods("PATCH")
	http.ListenAndServe(":8084", router)
}
