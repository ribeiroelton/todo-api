package storage

import "github.com/el7onr/go-todo/model"

type Repo interface {
	CreateToDo(*model.ToDo) (*model.ToDo, error)
	DeleteToDo(id int) error
	UpdateToDo(*model.ToDo) (*model.ToDo, error)
	ReadToDo(id int) (*model.ToDo, error)
	ListToDo() []*model.ToDo
}
