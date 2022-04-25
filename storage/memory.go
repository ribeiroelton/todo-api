package storage

import (
	"errors"
	"sync"

	"github.com/el7onr/go-todo/model"
)

type Database struct {
	storage map[int]*model.ToDo
	id      int
	mux     sync.Mutex
}

func NewDatabase() Repo {
	return &Database{
		storage: make(map[int]*model.ToDo),
		id:      0,
		mux:     sync.Mutex{},
	}
}

func (d *Database) CreateToDo(m *model.ToDo) (*model.ToDo, error) {
	d.mux.Lock()
	d.storage[d.id] = m
	d.mux.Unlock()
	m.Id = d.id
	d.id++

	return m, nil
}

func (d *Database) UpdateToDo(m *model.ToDo) (*model.ToDo, error) {
	d.mux.Lock()
	if _, ok := d.storage[m.Id]; ok {
		d.storage[m.Id] = m
		d.mux.Unlock()
		return m, nil

	}
	d.mux.Unlock()
	return nil, errors.New("id not found")
}

func (d *Database) DeleteToDo(id int) error {
	d.mux.Lock()
	if _, ok := d.storage[id]; ok {
		delete(d.storage, id)
		d.mux.Unlock()
		return nil

	}
	d.mux.Unlock()
	return errors.New("id not found")
}

func (d *Database) ReadToDo(id int) (*model.ToDo, error) {
	if _, ok := d.storage[id]; ok {
		return d.storage[id], nil
	}
	return nil, errors.New("id not found")
}

func (d *Database) ListToDo() []*model.ToDo {
	list := []*model.ToDo{}

	for _, v := range d.storage {
		list = append(list, v)
	}

	return list
}
