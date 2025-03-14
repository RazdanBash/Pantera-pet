package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"myProject/internal/database"
	"myProject/internal/handlers"
	"myProject/internal/taskService"
	"myProject/internal/userService"
	"myProject/internal/web/tasks"
	"myProject/internal/web/users"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&taskService.Task{})
	if err != nil {
		log.Fatalf("Error migrating tasks table: %v", err)
	}

	// Миграция таблицы пользователей
	err = database.DB.AutoMigrate(&userService.User{})
	if err != nil {
		log.Fatalf("Error migrating users table: %v", err)
	}

	tasksRepo := taskService.NewTaskRepository(database.DB)
	usersRepo := userService.NewUserRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)
	usersService := userService.NewUserService(usersRepo)

	tasksHandler := handlers.Handler{Service: tasksService}
	usersHandler := handlers.UserHandler{Service: usersService}

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictTasksHandler := tasks.NewStrictHandler(&tasksHandler, nil) //
	strictUsersHandler := users.NewStrictHandler(&usersHandler, nil)

	tasks.RegisterHandlers(e, strictTasksHandler)
	users.RegisterHandlers(e, strictUsersHandler)

	if err := e.Start(":8084"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
