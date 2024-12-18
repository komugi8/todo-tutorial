package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/komugi8/todo-tutorial/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
    var handler handler.Handler
    db, err := sqlx.Connect("mysql", "user:password@tcp(127.0.0.1:3306)/todo")
    if err != nil {
        panic(err)
    }
    handler.DB = db
	e := echo.New()
    e.Use(middleware.Logger())
	e.GET("/tasks", handler.GetTasks)
	e.GET("/tasks/:id", handler.GetTask)
    e.POST("/tasks", handler.CreateTask)
    e.PUT("/tasks/:id", handler.UpdateTask)
    e.DELETE("tasks/:id", handler.DeleteTask)
	e.Logger.Fatal(e.Start(":3030"))
}
