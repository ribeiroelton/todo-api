package api

import (
	"net/http"
	"strconv"

	"github.com/el7onr/go-todo/model"
	"github.com/el7onr/go-todo/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type apiHandler struct {
	server *echo.Echo
	db     storage.Repo
}

func NewApiHandler(e *echo.Echo, db storage.Repo) {

	cors := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS", "HEADER"},
	}

	e.Use(middleware.CORSWithConfig(cors))
	e.Use(middleware.Logger())

	a := &apiHandler{
		server: e,
		db:     db,
	}

	e.POST("/todos", a.createToDo)
	e.GET("/todos", a.listToDo)
	e.GET("/todos/:id", a.getToDo)
	e.DELETE("/todos/:id", a.deleteToDo)
	e.PUT("/todos/:id", a.updateToDo)

}

func (h *apiHandler) createToDo(c echo.Context) error {
	m := &model.ToDo{}

	err := c.Bind(m)
	if err != nil {
		c.JSON(http.StatusBadRequest, generateErrorResponse("invalid json", err))
		return err
	}

	r, err := h.db.CreateToDo(m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorResponse("error while writing data", err))
		return err
	}

	c.JSON(http.StatusOK, r)

	return nil
}

func (h *apiHandler) listToDo(c echo.Context) error {
	r := h.db.ListToDo()

	c.JSON(http.StatusOK, r)

	return nil
}

func (h *apiHandler) getToDo(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, generateErrorResponse("invalid parameter", err))
		return err
	}

	r, err := h.db.ReadToDo(id)
	if err != nil && err.Error() == "id not found" {
		c.NoContent(http.StatusNotFound)
		return err
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorResponse("eror while getting data", err))
		return err
	}

	c.JSON(http.StatusOK, r)

	return nil
}

func (h *apiHandler) deleteToDo(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, generateErrorResponse("invalid parameter", err))
		return err
	}

	err = h.db.DeleteToDo(id)
	if err != nil && err.Error() == "id not found" {
		c.NoContent(http.StatusNotFound)
		return err
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorResponse("error while deleting data", err))
		return err
	}

	c.NoContent(http.StatusOK)

	return nil

}

func (h *apiHandler) updateToDo(c echo.Context) error {
	m := &model.ToDo{}

	err := c.Bind(m)
	if err != nil {
		c.JSON(http.StatusBadRequest, generateErrorResponse("invalid json", err))
		return err
	}

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, generateErrorResponse("invalid parameter", err))
		return err
	}

	m.Id = id

	r, err := h.db.UpdateToDo(m)
	if err != nil && err.Error() == "id not found" {
		c.NoContent(http.StatusNotFound)
		return err
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorResponse("error while updating data", err))
		return err
	}

	c.JSON(http.StatusAccepted, r)
	return nil

}

func generateErrorResponse(m string, err error) *model.ErrorMessage {
	return &model.ErrorMessage{
		Message: m,
		Details: err.Error(),
	}
}
