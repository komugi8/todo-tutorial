package domain

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Tasks = []Task{
	{ID: 1, Name: "task1"},
	{ID: 2, Name: "task2"},
	{ID: 3, Name: "task3"},
	{ID: 4, Name: "task4"},
}
