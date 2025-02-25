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

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id

	err := h.Service.DeleteTask(taskID)
	if err != nil {
		if err == taskService.ErrTaskNotFound {
			return tasks.DeleteTasksId404JSONResponse{}, nil
		}
		return tasks.DeleteTasksId500JSONResponse{}, nil
	}

	return tasks.DeleteTasksId204Response{}, nil
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
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
	}

	return response, nil
}

func (h *Handler) GetTasks(ctx context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

//func (h *Handler) DeleteTaskId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
//	taskID := request.Id
//
//	err := h.Service.DeleteTask(taskID)
//	if err != nil {
//		if err == taskService.ErrTaskNotFound {
//			return tasks.DeleteTasksId404JSONResponse{}, nil
//		}
//		return tasks.DeleteTasksId500JSONResponse{}, nil
//	}
//
//	return tasks.DeleteTasksId204Response{}, nil
//}
//
//func (h *Handler) PatchTaskId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
//	taskID := request.Id
//	taskUpdate := request.Body
//
//	updatedTask := taskService.Task{
//		Model:  gorm.Model{},
//		Task:   *taskUpdate.Task,
//		IsDone: *taskUpdate.IsDone,
//	}
//
//	updatedTaskResult, err := h.Service.UpdateTaskByID(taskID, updatedTask)
//	if err != nil {
//		if errors.Is(err, taskService.ErrTaskNotFound) {
//			return tasks.PatchTasksId404JSONResponse{}, nil
//		}
//		return tasks.PatchTasksId500JSONResponse{}, nil
//	}
//
//	response := tasks.PatchTasksId200JSONResponse{
//		Id:     &updatedTaskResult.ID,
//		Task:   &updatedTaskResult.Task,
//		IsDone: &updatedTaskResult.IsDone,
//	}
//
//	return response, nil
//}
