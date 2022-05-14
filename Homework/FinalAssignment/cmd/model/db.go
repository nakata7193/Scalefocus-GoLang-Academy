package model

import "database/sql"

func DbInit() *Repository {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		panic(err)
	}
	repository := NewRepository(db)
	return repository
}