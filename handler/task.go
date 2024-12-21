package handler

import (
	"github.com/komugi8/todo-tutorial/domain/model"
	"github.com/komugi8/todo-tutorial/usecase"
	"github.com/labstack/echo/v4"
)

type TaskHandler interface {
	GetTasks(c echo.Context) error
	GetTask(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskHandler struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskHandler(taskUsecase usecase.TaskUsecase) TaskHandler {
	return &taskHandler{
		taskUsecase: taskUsecase,
	}
}

func (h *taskHandler) GetTasks(c echo.Context) error {
	tasks, err := h.taskUsecase.GetTasks()
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, tasks)
}

func (h *taskHandler) GetTask(c echo.Context) error {
	id := c.Param("id")
	task, err := h.taskUsecase.GetTask(id)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, task)
}

func (h *taskHandler) CreateTask(c echo.Context) error {
	task := new(model.Task)
	if err := c.Bind(task); err != nil {
		return c.JSON(400, err.Error())
	}
	createdTask, err := h.taskUsecase.CreateTask(*task)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(201, createdTask)
}

func (h *taskHandler) UpdateTask(c echo.Context) error {
	task := new(model.Task)
	if err := c.Bind(task); err != nil {
		return c.JSON(400, err.Error())
	}
	updatedTask, err := h.taskUsecase.UpdateTask(*task)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, updatedTask)
}

func (h *taskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	err := h.taskUsecase.DeleteTask(id)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.NoContent(204)
}