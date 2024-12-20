package handler

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/komugi8/todo-tutorial/domain"
	"github.com/labstack/echo/v4"
)


type Handler struct {
    DB *sqlx.DB
}

func (h *Handler) GetTasks(c echo.Context) error {
	tasks := []domain.Task{}
	err := h.DB.Select(&tasks, "SELECT * FROM tasks")
	if err != nil {
		log.Printf("Failed to get tasks: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to get tasks")
	}
	return c.JSON(http.StatusOK, &tasks)
}

func (h *Handler) GetTask(c echo.Context) error {
	id := c.Param("id")
	task := domain.Task{}
	err := h.DB.Get(&task, "SELECT * FROM tasks WHERE id = ?", id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Task not found")
	}
	return c.JSON(http.StatusOK, &task)
}

func (h *Handler) CreateTask(c echo.Context) error {
	task := domain.Task{}
    err := c.Bind(&task)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid request")
    }
	_, err = h.DB.Exec("INSERT INTO tasks (title, completed) VALUES (?, ?)", task.Title, task.Completed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create task")
	}
	return c.JSON(http.StatusCreated, &task)
}

func (h *Handler) UpdateTask(c echo.Context) error {
    id := c.Param("id")
    task := domain.Task{}
    err := c.Bind(&task)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid request")
    }
	_, err = h.DB.Exec("UPDATE tasks SET title = ?, completed = ? WHERE id = ?", task.Title, task.Completed, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update task")
	}
	return c.JSON(http.StatusOK, &task)
}

func (h *Handler) DeleteTask(c echo.Context) error {
    id := c.Param("id")
	_, err := h.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete task")
	}
	return c.NoContent(http.StatusNoContent)
}
