package repository

import "github.com/el7onr/go-todo/internal/model"

type Repo interface {
	CreateToDo(*model.ToDo) (*model.ToDo, error)
	DeleteToDo(id int) error
	UpdateToDo(*model.ToDo) (*model.ToDo, error)
	ReadToDo(id int) (*model.ToDo, error)
	ListTodo() ([]*model.ToDo, error)
}
