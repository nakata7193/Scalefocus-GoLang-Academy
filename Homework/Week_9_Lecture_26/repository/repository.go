package repository

import (
	"database/sql"
	"log"
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
func (repo *Repository) GetLastTimeStamp() time.Time {
	query := "SELECT s.timestamp FROM stories s ORDER BY s.timestamp DESC LIMIT 1"
	var timestamp time.Time
	if err := repo.db.QueryRow(query).Scan(&timestamp); err != nil {
		log.Print(err)
	}
	return timestamp
}

//getTopStories()
func (repo *Repository) GetStory() []stories.TopStory {
	query := "SELECT storyid,title,score FROM stories s"
	rows, err := repo.db.Query(query)
	if err != nil {
		log.Print(err)
	}

	defer rows.Close()
	stList := []stories.TopStory{}

	for {
		story := stories.TopStory{}
		if !rows.Next() {
			break
		}
		err := rows.Scan(&story.ID, &story.Title, &story.Score)
		if err != nil {
			log.Print(err)
		}
		stList = append(stList, story)
	}
	return stList
}

func (repo *Repository) SaveStories(stories []stories.TopStory) {
	query := "INSERT INTO stories (storyid, title, score) VALUES (?, ?, ?)"
	for _, story := range stories {
		repo.db.Exec(query, story.ID, story.Title, story.Score)
	}
}
