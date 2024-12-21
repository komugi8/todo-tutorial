package repository

import (
	"github.com/komugi8/todo-tutorial/domain/model"
)

type TaskRepository interface {
	GetTasks() ([]model.Task, error)
	GetTask(id int) (model.Task, error)
	CreateTask(task model.Task) (model.Task, error)
	UpdateTask(task model.Task) (model.Task, error)
	DeleteTask(id int) error
}