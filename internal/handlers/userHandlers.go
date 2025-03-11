package handlers

import (
	"myProject/internal/userService"
)

type Handler struct {
	Service *userService.UserService
}
