package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"myProject/internal/database"
	"myProject/internal/handlers"
	"myProject/internal/taskService"
	"myProject/internal/web/tasks"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&taskService.Task{})
	if err != nil {
		log.Fatalf("Error migrating tasks table: %v", err)
	}

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.Handler{Service: service}

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(&handler, nil) //
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8084"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
