package ports

import "github.com/ribeiroelton/todo-api/internal/domain"

type ToDoService interface {
	CreateToDo(*domain.ToDo) (*domain.ToDo, error)
	DeleteToDoById(id int) error
	UpdateToDo(*domain.ToDo) (*domain.ToDo, error)
	GetTodoById(id int) (*domain.ToDo, error)
	ListToDos() ([]*domain.ToDo, error)
}
