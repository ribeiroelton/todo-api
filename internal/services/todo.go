package services

import (
	"github.com/ribeiroelton/todo-api/internal/domain"
	"github.com/ribeiroelton/todo-api/internal/ports"
)

type ToDoService struct {
	db ports.ToDoRepository
}

func NewToDoService(db ports.ToDoRepository) *ToDoService {
	return &ToDoService{
		db: db,
	}
}

func (s *ToDoService) CreateToDo(m *domain.ToDo) (*domain.ToDo, error) {
	r, err := s.db.CreateToDo(m)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *ToDoService) DeleteToDoById(id int) error {
	err := s.db.DeleteToDoById(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *ToDoService) UpdateToDo(m *domain.ToDo) (*domain.ToDo, error) {
	r, err := s.db.UpdateToDo(m)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *ToDoService) GetTodoById(id int) (*domain.ToDo, error) {
	r, err := s.db.GetTodoById(id)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *ToDoService) ListToDos() ([]*domain.ToDo, error) {
	r, err := s.db.ListToDos()
	if err != nil {
		return nil, err
	}
	return r, err
}
