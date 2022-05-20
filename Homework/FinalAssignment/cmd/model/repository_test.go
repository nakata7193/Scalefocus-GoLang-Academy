package model

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

const (
	PragmaMeta      = "PRAGMA foreign_keys = ON"
	CreateListTable = "CREATE TABLE IF NOT EXISTS Lists(id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT)"
	CreateTaskTable = "CREATE TABLE IF NOT EXISTS Tasks(id INTEGER PRIMARY KEY AUTOINCREMENT,txt TEXT,list_id INTEGER,completed BOOLEAN NOT NULL ,FOREIGN KEY (list_id) REFERENCES Lists(id) ON DELETE CASCADE)"

	GetLists   = "SELECT * FROM Lists"
	CreateList = "INSERT INTO Lists (name) VALUES (?)"
	DeleteList = "DELETE FROM Lists WHERE id = (?)"

	GetTasks   = "SELECT * FROM Tasks WHERE list_id = (?)"
	CreateTask = "INSERT INTO Tasks (txt, list_id) VALUES (?, ?)"
	ToggleTask = "UPDATE Tasks SET completed = NOT completed WHERE id = (?)"
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

	_, err = mockDB.Exec(PragmaMeta)
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
	list := List{Name: "Test List", ID: 1}
	repo.CreateList(list.Name)

	rows, err := repo.db.Query(GetLists)
	if err != nil {
		t.Error(err)
	}

	for rows.Next() {
		list := List{}
		rows.Scan(&list.ID, &list.Name)

		if list.Name != "Test List" {
			t.Errorf("Expected name to be Test List, got %s", list.Name)
		}

		if list.ID != 1 {
			t.Errorf("Expected id to be 1, got %d", list.ID)
		}
	}
}

func TestDeleteList(t *testing.T) {
	var lists []List
	repo := mockDbRepo()
	list := List{Name: "Test List", ID: 1}
	list2 := List{Name: "Test List", ID: 2}
	repo.db.Exec(CreateList, list.Name)
	repo.db.Exec(CreateList, list2.Name)

	err := repo.DeleteList(list.ID)
	if err != nil {
		t.Error(err)
	}

	rows, err := repo.db.Query("SELECT * FROM Lists")
	if err != nil {
		t.Error(err)
	}

	for rows.Next() {
		list := List{}
		err := rows.Scan(&list.ID, &list.Name)
		if err != nil {
			t.Error(err)
		}
		lists = append(lists, list)

		if list.ID == 1 {
			t.Errorf("Expected id to be 2, got %d", list.ID)
		}
	}
}

//not working
func TestGetTasks(t *testing.T) {
	repo := mockDbRepo()
	list := List{Name: "Test List", ID: 1}
	repo.db.Exec(CreateList, list.Name)

	task := Task{Text: "Test Task", ListID: list.ID}
	repo.db.Exec(CreateTask, task.Text, task.ListID)

	tasks, err := repo.GetTasks(list.ID)
	if err != nil {
		t.Error(err)
	}

	//this needs to happen inside of rows.Next()
	if len(tasks) != 1 {
		t.Errorf("Expected 1 tasks, got %d", len(tasks))
	}
}

func TestCreateTask(t *testing.T) {
	repo := mockDbRepo()
	list := List{Name: "Test List", ID: 1}
	repo.db.Exec(CreateList, list.Name)

	task := Task{Text: "Test Task", ListID: list.ID}
	repo.db.Exec(CreateTask, task.Text, task.ListID)

	rows, err := repo.db.Query("SELECT * FROM Tasks")
	if err != nil {
		t.Error(err)
	}

	for rows.Next() {
		var tasks []Task
		rows.Scan(&tasks)

		if len(tasks) != 1 {
			t.Errorf("Expected 1 task, got %d", len(tasks))
		}
		if tasks[0].Text != "Test Task" {
			t.Errorf("Expected text to be Test Task, got %s", tasks[0].Text)
		}
		if tasks[0].ListID != 1 {
			t.Errorf("Expected listID to be 1, got %d", tasks[0].ListID)
		}
		if tasks[0].Completed != false {
			t.Errorf("Expected completed to be false, got %t", tasks[0].Completed)
		}
	}
}

func TestToggleTask(t *testing.T) {
	repo := mockDbRepo()
	list := List{Name: "Test List"}
	repo.db.Exec(CreateList, list.Name)

	task := Task{Text: "Test Task", ListID: list.ID}
	repo.db.Exec(CreateTask, task.Text, list.ID)

	err := repo.ToggleTask(task.ID)
	if err != nil {
		t.Error(err)
	}

	rows, err := repo.db.Query("SELECT * FROM Tasks")
	if err != nil {
		t.Error(err)
	}

	for rows.Next() {
		var task Task
		rows.Scan(&task)

		if task.Completed != true {
			t.Errorf("Expected completed to be true, got %t", task.Completed)
		}
	}
}

func TestDeleteTask(t *testing.T) {
	repo := mockDbRepo()

	list := List{Name: "Test List"}
	repo.db.Exec(CreateList, list.Name)
	task := Task{Text: "Test Task", ListID: list.ID}
	task2 := Task{Text: "Test Task 2", ListID: list.ID}
	repo.db.Exec(CreateTask, task2.Text, task2.ListID)
	repo.db.Exec(CreateList, task.Text, task.ListID)

	repo.DeleteTask(task.ID)

	rows, err := repo.db.Query("SELECT * FROM Tasks")
	if err != nil {
		t.Error(err)
	}

	for rows.Next() {
		var tasks []Task
		rows.Scan(&tasks)

		if len(tasks) != 1 {
			t.Errorf("Expected 1 task, got %d", len(tasks))
		}
	}
}

func TestCSVFile(t *testing.T) {
	repo := mockDbRepo()
	list := List{Name: "Test List"}
	repo.db.Exec(CreateList, list.Name)
	task := Task{Text: "Test Task", ListID: list.ID}
	repo.db.Exec(CreateTask, task.Text, task.ListID)

	file, err := repo.CSVExport()
	if err != nil {
		t.Error(err)
	}

	if file == nil {
		t.Error("Expected file to be not nil")
	}

	// content, err := ioutil.ReadFile("tasks.csv")
	// if err != nil {
	// 	t.Error(err)
	// }

	// if string(content) != "Test Task" {
	// 	t.Errorf("Expected content to be Test Task, got %s", string(content))
	// }
}
