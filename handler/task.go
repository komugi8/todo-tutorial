package handler

import (
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/komugi8/todo-tutorial/domain"
	"github.com/labstack/echo/v4"
)


type Handler struct {
    DB *sqlx.DB
}

func (h *Handler) GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, &domain.Tasks)
}

func (h *Handler) GetTask(c echo.Context) error {
	id := c.Param("id")
	for _, task := range domain.Tasks {
		if id == strconv.Itoa(task.ID) {
			return c.JSON(http.StatusOK, &task)
		}
	}
	return c.JSON(http.StatusNotFound, "Task not found")
}

func (h *Handler) CreateTask(c echo.Context) error {
    var task domain.Task
    err := c.Bind(&task)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid request")
    }
    domain.Tasks = append(domain.Tasks, task)
    return c.JSON(http.StatusCreated, &task)
}

func (h *Handler) UpdateTask(c echo.Context) error {
    id := c.Param("id")
    var task domain.Task
    err := c.Bind(&task)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid request")
    }
    for i, t := range domain.Tasks {
        if id == strconv.Itoa(t.ID) {
            domain.Tasks[i] = task
            return c.JSON(http.StatusOK, &task)
        }
    }
    return c.JSON(http.StatusNotFound, "Task not found")
}

func (h *Handler) DeleteTask(c echo.Context) error {
    id := c.Param("id")
    for i, task := range domain.Tasks {
        if id == strconv.Itoa(task.ID) {
            domain.Tasks = append(domain.Tasks[:i], domain.Tasks[i+1:]...)
            return c.JSON(http.StatusOK, &task)
        }
    }
    return c.JSON(http.StatusNotFound, "Task not found")
}
