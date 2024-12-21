package cmd

import (
	"github.com/komugi8/todo-tutorial/handler"
	"github.com/komugi8/todo-tutorial/infrastructure"
	"github.com/komugi8/todo-tutorial/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunRouter() {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}
	db, err := cfg.GetDB()
	if err != nil {
		panic(err)
	}

	taskRepository := infrastructure.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/healthcheck", handler.Healthcheck)
	e.GET("/tasks", taskHandler.GetTasks)
	e.GET("/tasks/:id", taskHandler.GetTask)
	e.POST("/tasks", taskHandler.CreateTask)
	e.PUT("/tasks", taskHandler.UpdateTask)
	e.DELETE("/tasks/:id", taskHandler.DeleteTask)
	e.Logger.Fatal(e.Start(":3030"))
}