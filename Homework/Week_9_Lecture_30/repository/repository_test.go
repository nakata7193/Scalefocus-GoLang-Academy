package repository

import (
	"database/sql"
	"testing"
	"time"

	_ "modernc.org/sqlite"
)

const (
	createStoryTable = "CREATE TABLE IF NOT EXISTS stories(storyid INT PRIMARY KEY,title TEXT,score INT,timestamp DATETIME DEFAULT CURRENT_TIMESTAMP)"
	insertStory      = "INSERT INTO stories (storyid, title, score) VALUES (?, ?, ?,?)"
)

func TestLastStoryTimeStamp(t *testing.T) {
	mockDB, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open mock DB")
	}
	_, err = mockDB.Exec(createStoryTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}

	repo := NewRepository(mockDB)
	result := repo.GetLastTimeStamp()
	if result == time.Now() {
		t.Errorf("Failed to create initial condition")
	}
	wantedTime := time.Now().Add(time.Hour)

	mockDB.Exec(insertStory, 1, "test", 1, wantedTime)
	mockDB.Exec(insertStory, 2, "test1", 2, time.Now().Add(-time.Hour))
	result = repo.GetLastTimeStamp()

	if result == wantedTime {
		t.Errorf("Failed to get last timestamp")
	}
}
