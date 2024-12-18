package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/komugi8/todo-tutorial/cmd"
	"github.com/komugi8/todo-tutorial/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
    cfg, err := cmd.NewConfig()
    if err != nil {
        panic(err)
    }
    db, err := cfg.GetDB()
    if err != nil {
        panic(err)
    }

    handler := handler.Handler{}
    handler.DB = db

	e := echo.New()
    e.Use(middleware.Logger())
    e.GET("/healthcheck", handler.Healthcheck)
	e.GET("/tasks", handler.GetTasks)
	e.GET("/tasks/:id", handler.GetTask)
    e.POST("/tasks", handler.CreateTask)
    e.PUT("/tasks/:id", handler.UpdateTask)
    e.DELETE("tasks/:id", handler.DeleteTask)
	e.Logger.Fatal(e.Start(":3030"))
}
