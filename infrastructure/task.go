package infrastructure

import (
	"github.com/jmoiron/sqlx"
	"github.com/komugi8/todo-tutorial/domain/model"
	"github.com/komugi8/todo-tutorial/domain/repository"
)

type TaskRepositoryImpl struct{
	DB *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) repository.TaskRepository {
	return &TaskRepositoryImpl{DB: db}
}

func (r *TaskRepositoryImpl) GetTasks() ([]model.Task, error) {
	tasks := []model.Task{}
	err := r.DB.Select(&tasks, "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepositoryImpl) GetTask(id int) (model.Task, error) {
	task := model.Task{}
	err := r.DB.Get(&task, "SELECT * FROM tasks WHERE id = ?", id)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (r *TaskRepositoryImpl) CreateTask(task model.Task) (model.Task, error) {
	_, err := r.DB.Exec("INSERT INTO tasks (title, completed) VALUES (?, ?)", task.Title, task.Completed)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (r *TaskRepositoryImpl) UpdateTask(task model.Task) (model.Task, error) {
	_, err := r.DB.Exec("UPDATE tasks SET title = ?, completed = ? WHERE id = ?", task.Title, task.Completed, task.ID)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (r *TaskRepositoryImpl) DeleteTask(id int) error {
	_, err := r.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}