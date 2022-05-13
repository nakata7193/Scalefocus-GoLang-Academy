package controllers

type List interface {
	GetLists() ([]List, error)
	CreateList(listName string) (List, error)
	DeleteList(listID int) error
}


