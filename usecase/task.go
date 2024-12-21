package usecase

import (
	"strconv"

	"github.com/komugi8/todo-tutorial/domain/model"
	"github.com/komugi8/todo-tutorial/domain/repository"
)

type TaskUsecase interface {
	GetTasks() ([]model.Task, error)
	GetTask(id string) (model.Task, error)
	CreateTask(task model.Task) (model.Task, error)
	UpdateTask(task model.Task) (model.Task, error)
	DeleteTask(id string) error
}

type taskUsecase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUsecase(taskRepository repository.TaskRepository) TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
	}
}

func (u *taskUsecase) GetTasks() ([]model.Task, error) {
	return u.taskRepository.GetTasks()
}

func (u *taskUsecase) GetTask(id string) (model.Task, error) {
	newID, err := strconv.Atoi(id)
	if err != nil {
		return model.Task{}, err
	}
	return u.taskRepository.GetTask(newID)
}

func (u *taskUsecase) CreateTask(task model.Task) (model.Task, error) {
	return u.taskRepository.CreateTask(task)
}

func (u *taskUsecase) UpdateTask(task model.Task) (model.Task, error) {
	return u.taskRepository.UpdateTask(task)
}

func (u *taskUsecase) DeleteTask(id string) error {
	newID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return u.taskRepository.DeleteTask(newID)
}