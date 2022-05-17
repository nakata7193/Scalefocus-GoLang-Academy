package utils

import (
	"database/sql"
	"final/cmd/model"

	_ "modernc.org/sqlite"
)

func DbInit() *model.Repository {
	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		panic(err)
	}
	repository := model.NewRepository(db)
	return repository
}

