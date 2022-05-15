package model

import (
	"database/sql"
	"reflect"
	"testing"

	_ "modernc.org/sqlite"
)

type FakeStorage struct {
}

const (
	CreateListTable = "CREATE TABLE IF NOT EXISTS Lists(id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT)"
	CreateTaskTable = "CREATE TABLE IF NOT EXISTS Tasks(id INTEGER PRIMARY KEY AUTOINCREMENT,txt TEXT,list_id INTEGER,completed BOOLEAN NOT NULL DEFAULT 0,FOREIGN KEY (list_id) REFERENCES Lists(id) ON DELETE CASCADE)"

	CreateList = "INSERT INTO Lists (name) VALUES (?)"
	DeleteList = "DELETE FROM Lists WHERE id = (?)"

	CreateTask = "INSERT INTO Tasks (txt, list_id) VALUES (?, ?)"
	ToggleTask = "SELECT id, txt, completed FROM Tasks WHERE id = (?)"
	DeleteTask = "DELETE FROM Tasks WHERE id = (?)"
)

func mockDbRepo() *Repository {
	mockDB, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		panic(err)
	}

	_, err = mockDB.Exec(CreateListTable)
	if err != nil {
		panic(err)
	}

	_, err = mockDB.Exec(CreateTaskTable)
	if err != nil {
		panic(err)
	}

	repo := NewRepository(mockDB)
	return repo
}

func TestGetLists(t *testing.T) {
	repo := mockDbRepo()
	list := List{Name: "Test List"}
	repo.db.Exec(CreateList, list.Name)

	lists, err := repo.GetLists()
	if err != nil {
		t.Error(err)
	}

	if len(lists) != 1 {
		t.Errorf("Expected 1 list, got %d", len(lists))
	}

}

func TestCreateList(t *testing.T) {
	repo := mockDbRepo()
	list := List{Name: "Test List"}
	repo.CreateList(list)
	lists, err := repo.GetLists()
	if err != nil {
		t.Error(err)
	}

	if len(lists) != 1 {
		t.Errorf("Expected 1 list, got %d", len(lists))
	}

}

func TestDeleteList(t *testing.T) {
	repo := mockDbRepo()
	list := List{Name: "Test List"}
	repo.CreateList(list)
	repo.DeleteList(list)
	lists, err := repo.GetLists()
	if err != nil {
		t.Error(err)
	}

	if len(lists) != 0 {
		t.Errorf("Expected 0 lists, got %d", len(lists))
	}
}

func TestGetTasks(t *testing.T) {
	repo := mockDbRepo()
	list := List{Name: "Test List"}
	repo.db.Exec(CreateList, list.Name)

	task := Task{Text: "Test Task", ListID: list.ID}
	repo.db.Exec(CreateTask, task.Text, task.ListID)

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
	repo.db.Exec(CreateList, list.Name)

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
	repo.db.Exec(CreateList, list.Name)

	task := Task{Text: "Test Task", ListID: list.ID}
	repo.db.Exec(CreateTask, task.Text, task.ListID)

	completedTask, err := repo.ToggleTask(task)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(completedTask, task) {
		t.Errorf("Expected %v, got %v", task, completedTask)
	}
}

func TestDeleteTask(t *testing.T) {
	repo := mockDbRepo()

	list := List{Name: "Test List"}
	repo.db.Exec(CreateList, list.Name)
	task := Task{Text: "Test Task", ListID: list.ID}
	repo.db.Exec(CreateList, task.Text, task.ListID)

	repo.DeleteTask(task)

	tasks, err := repo.GetTasks(list)
	if err != nil {
		t.Error(err)
	}

	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(tasks))
	}
}
