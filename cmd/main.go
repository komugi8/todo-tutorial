package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var tasks = []Task{
	{ID: 1, Name: "task1"},
	{ID: 2, Name: "task2"},
	{ID: 3, Name: "task3"},
	{ID: 4, Name: "task4"},
}

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, &tasks)
}

func getTask(c echo.Context) error {
	id := c.Param("id")
	for _, task := range tasks {
		if id == strconv.Itoa(task.ID) {
			return c.JSON(http.StatusOK, &task)
		}
	}
	return c.JSON(http.StatusNotFound, "Task not found")
}

func createTask(c echo.Context) error {
    var task Task
    err := c.Bind(&task)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid request")
    }
    tasks = append(tasks, task)
    return c.JSON(http.StatusCreated, &task)
}

func updateTask(c echo.Context) error {
    id := c.Param("id")
    var task Task
    err := c.Bind(&task)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid request")
    }
    for i, t := range tasks {
        if id == strconv.Itoa(t.ID) {
            tasks[i] = task
            return c.JSON(http.StatusOK, &task)
        }
    }
    return c.JSON(http.StatusNotFound, "Task not found")
}

func deleteTask(c echo.Context) error {
    id := c.Param("id")
    for i, task := range tasks {
        if id == strconv.Itoa(task.ID) {
            tasks = append(tasks[:i], tasks[i+1:]...)
            return c.JSON(http.StatusOK, &task)
        }
    }
    return c.JSON(http.StatusNotFound, "Task not found")
}

func main() {
	e := echo.New()
    e.Use(middleware.Logger())
	e.GET("/tasks", getTasks)
	e.GET("/tasks/:id", getTask)
    e.POST("/tasks", createTask)
    e.PUT("/tasks/:id", updateTask)
    e.DELETE("tasks/:id", deleteTask)
	e.Logger.Fatal(e.Start(":3030"))
}
