package handlers

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"myProject/internal/taskService"
	"myProject/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id

	err := h.Service.DeleteTask(taskID)
	if err != nil {
		if errors.Is(err, taskService.ErrTaskNotFound) {
			return tasks.DeleteTasksId404JSONResponse{}, nil
		}
		return tasks.DeleteTasksId500JSONResponse{}, nil
	}

	return tasks.DeleteTasksId204Response{}, nil
}

func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := request.Id
	taskUpdate := request.Body

	updatedTask := taskService.Task{
		Model:  gorm.Model{},
		Task:   *taskUpdate.Task,
		IsDone: *taskUpdate.IsDone,
	}

	updatedTaskResult, err := h.Service.UpdateTaskByID(taskID, updatedTask)
	if err != nil {
		if errors.Is(err, taskService.ErrTaskNotFound) {
			return tasks.PatchTasksId404JSONResponse{}, nil
		}
		return tasks.PatchTasksId500JSONResponse{}, nil
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTaskResult.ID,
		Task:   &updatedTaskResult.Task,
		IsDone: &updatedTaskResult.IsDone,
		UserId: &updatedTaskResult.UserID,
	}

	return response, nil
}

func (h *Handler) GetTasksId(_ context.Context, request tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	userID := request.Id
	allTasks, err := h.Service.GetTaskByUserId(userID)
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasksId200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) PostTasksId(_ context.Context, request tasks.PostTasksIdRequestObject) (tasks.PostTasksIdResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	UserID := request.Id
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(UserID, taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasksId201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}
	// Просто возвращаем респонс!
	return response, nil
}
