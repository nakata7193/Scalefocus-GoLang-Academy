package model

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Repository struct {
	db *sql.DB
}

type TaskOperations interface {
	GetTasks(list List) ([]Task, error)
	CreateTask(task Task, list List) (Task, error)
	ToggleTask(task Task) (Task, error)
	DeleteTask(task Task) error
}

type ListOperations interface {
	GetLists() ([]List, error)
	CreateList(list List) (List, error)
	DeleteList(list List) error
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetTasks(list List) ([]Task, error) {
	query := "SELECT txt, completed from Tasks WHERE list_id = ?"
	rows, err := r.db.Query(query, list.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []Task{}
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Text, &task.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *Repository) CreateTask(task Task, list List) (Task, error) {
	query := "INSERT INTO Tasks (txt, list_id) VALUES (?, ?)"
	err := r.db.QueryRow(query, task.Text, list.ID).Scan(&task.ID, &task.Text, &task.ListID, &task.Completed)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *Repository) ToggleTask(task Task) (Task, error) {
	query := "UPDATE Tasks SET completed = NOT completed WHERE id = (?)"
	_, err := r.db.Exec(query, task.ID)
	if err != nil {
		return task, err
	}

	query = "SELECT id, txt, completed FROM Tasks WHERE id = (?)"
	resultTask := r.db.Query(task.ID).Scan(&task.ID, &task.Text, &task.ListID, &task.Completed)

	return task, nil
}

func (r *Repository) DeleteTask(task Task) error {
	_, err := r.db.Exec("DELETE FROM Tasks WHERE id = ?)", task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetLists() ([]List, error) {
	query := "SELECT * FROM Lists"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lists := []List{}
	for rows.Next() {
		var list List
		err := rows.Scan(&list.ID, &list.Name)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}

func (r *Repository) CreateList(list List) (List, error) {
	query := "INSERT INTO Lists (name) VALUES (?)"
	err := r.db.QueryRow(query, list.Name).Scan(&list.ID, &list.Name)
	if err != nil {
		return list, err
	}

	return list, nil
}

func (r *Repository) DeleteList(list List) error {
	_, err := r.db.Exec("DELETE FROM Lists WHERE id = ?", list.ID)
	if err != nil {
		return err
	}
	return nil
}
