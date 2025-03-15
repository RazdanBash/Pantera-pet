package userService

import (
	"gorm.io/gorm"
	"myProject/internal/taskService"
)

type User struct {
	gorm.Model
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []taskService.Task `json:"tasks" gorm:"foreignKey:UserID"`
}
