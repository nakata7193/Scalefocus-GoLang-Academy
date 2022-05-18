package model

type Task struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	ListID    int    `json:"listId"`
	Completed bool   `json:"completed"`
}

type List struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}