package model

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Repository struct {
	db *sql.DB
}

type TaskOperations interface {
	GetTasks(listID int) ([]Task, error)
	CreateTask(listID int, taskText string) error
	ToggleTask(taskID int) error
	DeleteTask(taskID int) error
}

type ListOperations interface {
	GetLists() ([]List, error)
	CreateList(listName string) error
	DeleteList(listID int) error
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetTasks(listID int) ([]Task, error) {
	tasks := []Task{}
	query := "SELECT * from Tasks WHERE list_id = (?)"
	rows, err := r.db.Query(query, listID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		task := Task{}
		err := rows.Scan(&task.ID, &task.Text, &task.ListID, &task.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *Repository) CreateTask(listID int, taskText string) error {
	query := "INSERT INTO Tasks (txt, list_id) VALUES (?, ?)"
	_, err := r.db.Exec(query, taskText, listID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) ToggleTask(taskID int) error {
	query := "UPDATE Tasks SET completed = NOT completed WHERE id = (?)"
	_, err := r.db.Exec(query, taskID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteTask(taskID int) error {
	_, err := r.db.Exec("DELETE FROM Tasks WHERE id = ?)", taskID)
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

func (r *Repository) CreateList(listName string) error {
	query := "INSERT INTO Lists (name) VALUES (?)"
	_, err := r.db.Exec(query, listName)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteList(listID int) error {
	_, err := r.db.Exec("DELETE FROM Lists WHERE id = (?)", listID)
	if err != nil {
		return err
	}
	return nil
}
