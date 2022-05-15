package model

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

type FakeStorage struct {
}

const (
	createListTable = "CREATE TABLE IF NOT EXISTS Lists(id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT)"
	createTaskTable = "CREATE TABLE IF NOT EXISTS Tasks(id INTEGER PRIMARY KEY AUTOINCREMENT,txt TEXT,list_id INTEGER,completed BOOLEAN NOT NULL DEFAULT 0,FOREIGN KEY (list_id) REFERENCES Lists(id) ON DELETE CASCADE)"
	
	getListsQuery   = "SELECT name from Lists"
	createListQuery = "INSERT INTO Lists (name) VALUES (?)"
	deleteListQuery = "DELETE FROM Lists WHERE id = ?"
	
	getTasksQuery   = "SELECT txt, completed from Tasks WHERE list_id = ?"
	createTaskQuery = "INSERT INTO Tasks (txt, list_id) VALUES (?, ?)"
	toggleTaskQuery = "UPDATE Tasks SET completed = NOT completed WHERE id = ?"
	deleteTaskQuery = "DELETE FROM Tasks WHERE id = ?"
)

func mockDbRepo() *Repository {
	mockDB, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		panic(err)
	}

	_, err = mockDB.Exec(createListTable)
	if err != nil {
		panic(err)
	}

	_, err = mockDB.Exec(createTaskTable)
	if err != nil {
		panic(err)
	}

	repo := NewRepository(mockDB)
	return repo
}

func TestGetTasks(t *testing.T) {
	repo := mockDbRepo()
	list := List{Name: "Test List"}
	repo.CreateList(list)
	task := Task{Text: "Test Task", ListID: list.ID}
	repo.CreateTask(task, list)
	tasks, err := repo.GetTasks(list)
	if err != nil {
		t.Error(err)
	}
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
}

func TestCreateTask(t *testing.T) {
	repo := mockDbRepo()
	list := List{Name: "Test List"}
	repo.CreateList(list)
	task := Task{Text: "Test Task", ListID: list.ID}
	repo.CreateTask(task, list)
	tasks, err := repo.GetTasks(list)
	if err != nil {
		t.Error(err)
	}
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
}

func TestToggleTask(t *testing.T) {
	repo := mockDbRepo()
	list := List{Name: "Test List"}
	repo.CreateList(list)
	task := Task{Text: "Test Task", ListID: list.ID}
	repo.CreateTask(task, list)
	repo.ToggleTask(task)
	
	if task.Completed != true {
		t.Errorf("Expected task to be completed, got %v", task.Completed)
	}
}

func TestDeleteTask(t *testing.T) {
	repo := mockDbRepo()
	list := List{Name: "Test List"}
	repo.CreateList(list)
	task := Task{Text: "Test Task", ListID: list.ID}
	repo.CreateTask(task, list)
	repo.DeleteTask(task)
	if list.
}


func TestGetLists(t *testing.T) {
	t.Skip("Not implemented")
}

func TestCreateList(t *testing.T) {
	t.Skip("Not implemented")
}

func TestDeleteList(t *testing.T) {
	t.Skip("Not implemented")
}
