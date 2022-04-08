package storage_test

import (
	"testing"

	"github.com/el7onr/go-todo/model"
	"github.com/el7onr/go-todo/storage"
	"github.com/stretchr/testify/assert"
)

func TestCreateToDo(t *testing.T) {
	db := storage.NewDatabase()
	r, err := db.CreateToDo(&model.ToDo{})
	assert.Nil(t, err)
	assert.NotNil(t, r)
}

func TestDeleteToDo_ok(t *testing.T) {
	db := storage.NewDatabase()
	db.CreateToDo(&model.ToDo{})

	err := db.DeleteToDo(0)
	assert.Nil(t, err)
}

func TestDeleteToDo_notfound(t *testing.T) {
	db := storage.NewDatabase()
	err := db.DeleteToDo(0)
	assert.Contains(t, err.Error(), "not found")
}

func TestListToDo_one(t *testing.T) {
	db := storage.NewDatabase()
	db.CreateToDo(&model.ToDo{})
	r := db.ListToDo()
	assert.Len(t, r, 1)
}

func TestListToDo_zero(t *testing.T) {
	db := storage.NewDatabase()
	r := db.ListToDo()
	assert.Len(t, r, 0)
}

func TestGetTodo_ok(t *testing.T) {
	db := storage.NewDatabase()
	db.CreateToDo(&model.ToDo{})
	m, err := db.ReadToDo(0)
	assert.Nil(t, err)
	assert.NotNil(t, m)
}

func TestGetTodo_not_found(t *testing.T) {
	db := storage.NewDatabase()
	m, err := db.ReadToDo(0)
	assert.Contains(t, err.Error(), "not found")
	assert.Nil(t, m)
}

func TestUpdateTodo_ok(t *testing.T) {
	db := storage.NewDatabase()
	db.CreateToDo(&model.ToDo{Title: "task1"})

	m, err := db.UpdateToDo(&model.ToDo{Title: "task2"})
	assert.Contains(t, m.Title, "task2")
	assert.Nil(t, err)
}

func TestUpdateTodo_not_found(t *testing.T) {
	db := storage.NewDatabase()

	m, err := db.UpdateToDo(&model.ToDo{Id: 1, Title: "task2"})
	assert.Contains(t, err.Error(), "not found")
	assert.Nil(t, m)
}
