package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ribeiroelton/todo-api/internal/domain"
	"github.com/ribeiroelton/todo-api/internal/ports"
)

type ToDoHandler struct {
	service ports.ToDoService
}

type ErrorMessage struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func NewToDoHandler(s ports.ToDoService) *ToDoHandler {
	return &ToDoHandler{
		service: s,
	}
}

func (a *ToDoHandler) CreateToDo(c echo.Context) error {
	m := &domain.ToDo{}

	err := c.Bind(m)
	if err != nil {
		c.JSON(http.StatusBadRequest, genError("invalid json", err))
		return err
	}

	r, err := a.service.CreateToDo(m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, genError("error while writing data", err))
		return err
	}

	c.JSON(http.StatusOK, r)

	return nil
}

func (a *ToDoHandler) ListToDos(c echo.Context) error {
	r, err := a.service.ListToDos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, r)

	return nil
}

func (a *ToDoHandler) GetToDo(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, genError("invalid parameter", err))
		return err
	}

	r, err := a.service.GetTodoById(id)
	if err != nil && err.Error() == "id not found" {
		c.NoContent(http.StatusNotFound)
		return err
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, genError("error while getting data", err))
		return err
	}

	c.JSON(http.StatusOK, r)

	return nil
}

func (a *ToDoHandler) DeleteToDo(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, genError("invalid parameter", err))
		return err
	}

	err = a.service.DeleteToDoById(id)
	if err != nil && err.Error() == "id not found" {
		c.NoContent(http.StatusNotFound)
		return err
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, genError("error while deleting data", err))
		return err
	}

	c.NoContent(http.StatusOK)

	return nil

}

func (a *ToDoHandler) UpdateToDo(c echo.Context) error {
	m := &domain.ToDo{}

	err := c.Bind(m)
	if err != nil {
		c.JSON(http.StatusBadRequest, genError("invalid json", err))
		return err
	}

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, genError("invalid parameter", err))
		return err
	}

	m.Id = id

	r, err := a.service.UpdateToDo(m)
	if err != nil && err.Error() == "id not found" {
		c.NoContent(http.StatusNotFound)
		return err
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, genError("error while updating data", err))
		return err
	}

	c.JSON(http.StatusAccepted, r)
	return nil

}

func genError(m string, err error) *ErrorMessage {
	return &ErrorMessage{
		Message: m,
		Details: err.Error(),
	}
}
