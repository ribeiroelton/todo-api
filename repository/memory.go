package repository

import (
	"errors"
	"sync"

	"github.com/ribeiroelton/todo-api/internal/domain"
)

type MemoryDB struct {
	storage map[int]*domain.ToDo
	id      int
	mux     sync.Mutex
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		storage: make(map[int]*domain.ToDo),
		id:      0,
		mux:     sync.Mutex{},
	}
}

func (d *MemoryDB) CreateToDo(m *domain.ToDo) (*domain.ToDo, error) {
	d.mux.Lock()
	d.storage[d.id] = m
	d.mux.Unlock()
	m.Id = d.id
	d.id++

	return m, nil
}

func (d *MemoryDB) UpdateToDo(m *domain.ToDo) (*domain.ToDo, error) {
	d.mux.Lock()
	if _, ok := d.storage[m.Id]; ok {
		d.storage[m.Id] = m
		d.mux.Unlock()
		return m, nil

	}
	d.mux.Unlock()
	return nil, errors.New("id not found")
}

func (d *MemoryDB) DeleteToDoById(id int) error {
	d.mux.Lock()
	if _, ok := d.storage[id]; ok {
		delete(d.storage, id)
		d.mux.Unlock()
		return nil

	}
	d.mux.Unlock()
	return errors.New("id not found")
}

func (d *MemoryDB) GetTodoById(id int) (*domain.ToDo, error) {
	if _, ok := d.storage[id]; ok {
		return d.storage[id], nil
	}
	return nil, errors.New("id not found")
}

func (d *MemoryDB) ListToDos() ([]*domain.ToDo, error) {
	list := []*domain.ToDo{}

	for _, v := range d.storage {
		list = append(list, v)
	}

	return list, nil
}
