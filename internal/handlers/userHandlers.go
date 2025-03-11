package handlers

import (
	"context"
	"errors"
	"myProject/internal/userService"
	"myProject/internal/web/users"
)

type userHandler struct {
	Service *userService.UserService
}

func (h *userHandler) DeleteUserByID(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userID := request.Id
	err := h.Service.DeleteUserByID(userID)
	if err != nil {
		if errors.Is(err, userService.ErrUserNotFound) {
			return users.DeleteUsersId404JSONResponse{}, nil
		}
		return users.DeleteUsersId500JSONResponse{}, nil
	}

	return users.DeleteUsersId204Response{}, nil
}

func (h *userHandler) GetAllUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	AllUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range AllUsers {
		user := users.User{
			Id:       &usr.ID,
			Username: &usr.Username,
			Password: &usr.Password,
			Email:    &usr.Email,
		}
		response = append(response, user)
	}
	return response, nil
}
