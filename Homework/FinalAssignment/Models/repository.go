package repository

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

//get all tasks from a list
func (r *Repository) GetTasks(ListID int) ([]Task, error) {
	query := "SELECT txt, completed from Tasks WHERE list_id = ?"
	rows, err := r.db.Query(query, id)
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

func (r *Repository) CreateTask(text string, listID int) (Task, error) {
	var task Task
	query := "INSERT INTO Tasks (txt, list_id) VALUES (?, ?)"
	err := r.db.QueryRow(query, text, listID).Scan(&task.ID, &task.Text, &task.ListID, &task.Completed)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *Repository) ToggleTask(listID int, taskID int) (Task, error) {
	var task Task
	query := "UPDATE Tasks SET completed = NOT completed WHERE (list_id = ? AND id = ?)"
	_, err := r.db.Exec(query, listID, taskID)
	if err != nil {
		return task, err
	}

	query = "SELECT id, txt, completed FROM Tasks WHERE (list_id = ? AND id = ?)"
	err = r.db.QueryRow(query, listID, taskID).Scan(&task.ID, &task.Text, &listID, &task.Completed)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *Repository) DeleteTask(listID int, taskID int) error {
	_, err := r.db.Exec("DELETE FROM Tasks WHERE (list_id = ? AND id = ?)", listID, taskID)
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

func (r *Repository) CreateList(listName string) (List, error) {
	var list List
	query := "INSERT INTO Lists (name) VALUES (?)"
	err := r.db.QueryRow(query, listName).Scan(&list.ID, &list.Name)
	if err != nil {
		return list, err
	}

	return list, nil
}

func (r *Repository) DeleteList(listID int) error {
	_, err := r.db.Exec("DELETE FROM Lists WHERE id = ?", listID)
	if err != nil {
		return err
	}
	return nil
}
