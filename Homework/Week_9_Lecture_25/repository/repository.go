package repository

import (
	"database/sql"
	"time"

	stories "github.com/nakata7193/story"
)

type Repository struct {
	db *sql.DB
}

//new repository(db)
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

//getLastTimeStamp()
func (repo *Repository)GetLastTimeStamp() time.Time {
	return time.Now()
}

//getTopStories()
func (repo *Repository)GetStory() []stories.TopStory {
	return []stories.TopStory{}
}
