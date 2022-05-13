package handlers

type List interface {
	GetLists() ([]List, error)
	CreateList(listName string) (List, error)
	DeleteList(listID int) error
}

type Task interface {
	GetTasks(ListID int) ([]Task, error)
	CreateTask(text string, listID int) (Task, error)
	ToggleTask(listID int, taskID int) (Task, error)
	DeleteTask(listID int, taskID int) error
}
