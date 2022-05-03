package repository

import (
	"os"
	"testing"

	"github.com/ribeiroelton/todo-api/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
func TestCreateToDo_ShouldCreate(t *testing.T) {
	db := NewMemoryDB()
	r, err := db.CreateToDo(&domain.ToDo{})
	assert.Nil(t, err)
	assert.NotNil(t, r)
}

func TestDeleteToDo_ShouldDelete(t *testing.T) {
	db := NewMemoryDB()
	db.CreateToDo(&domain.ToDo{})
	err := db.DeleteToDoById(0)
	assert.Nil(t, err)
}

func TestDeleteToDo_ShouldReturnNotFound(t *testing.T) {
	db := NewMemoryDB()
	err := db.DeleteToDoById(0)
	assert.Contains(t, err.Error(), "not found")
}

func TestListToDo_ShouldReturnOne(t *testing.T) {
	db := NewMemoryDB()
	db.CreateToDo(&domain.ToDo{})
	r, _ := db.ListToDos()
	assert.Len(t, r, 1)
}

func TestListToDo_ShouldReturnZero(t *testing.T) {
	db := NewMemoryDB()
	r, _ := db.ListToDos()
	assert.Len(t, r, 0)
}

func TestGetTodo_ShouldGetOne(t *testing.T) {
	db := NewMemoryDB()
	db.CreateToDo(&domain.ToDo{})
	m, err := db.GetTodoById(0)
	assert.Nil(t, err)
	assert.NotNil(t, m)
}

func TestGetTodo_ShouldReturnNotFound(t *testing.T) {
	db := NewMemoryDB()
	m, err := db.GetTodoById(0)
	assert.Contains(t, err.Error(), "not found")
	assert.Nil(t, m)
}

func TestUpdateTodo_ShouldUpdate(t *testing.T) {
	db := NewMemoryDB()
	db.CreateToDo(&domain.ToDo{Title: "task1"})

	m, err := db.UpdateToDo(&domain.ToDo{Title: "task2"})
	assert.Contains(t, m.Title, "task2")
	assert.Nil(t, err)
}

func TestUpdateTodo_ShouldReturnNotFound(t *testing.T) {
	db := NewMemoryDB()
	m, err := db.UpdateToDo(&domain.ToDo{Id: 1, Title: "task2"})
	assert.Contains(t, err.Error(), "not found")
	assert.Nil(t, m)
}
