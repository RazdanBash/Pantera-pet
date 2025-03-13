package handlers

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"myProject/internal/userService"
	"myProject/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func (h *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
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

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
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

func (h *UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userID := request.Id
	updateUser := request.Body

	updatedUser := userService.User{
		Model:    gorm.Model{},
		Username: *updateUser.Username,
		Password: *updateUser.Password,
		Email:    *updateUser.Email,
	}

	updatedUserResult, err := h.Service.UpdateUserByID(userID, updatedUser)
	if err != nil {
		if errors.Is(err, userService.ErrUserNotFound) {
			return users.PatchUsersId404JSONResponse{}, nil
		}

		return users.PatchUsersId500JSONResponse{}, nil
	}
	response := users.PatchUsersId200JSONResponse{
		Id:       &updatedUserResult.ID,
		Email:    &updatedUserResult.Email,
		Username: &updatedUserResult.Username,
		Password: &updatedUserResult.Password,
	}
	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := userService.User{
		Model:    gorm.Model{},
		Username: *userRequest.Username,
		Password: *userRequest.Password,
		Email:    *userRequest.Email,
	}

	createdUser, err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Username: &createdUser.Username,
		Password: &createdUser.Password,
		Email:    &createdUser.Email,
	}
	return response, nil
}
