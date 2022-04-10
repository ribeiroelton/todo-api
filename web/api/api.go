package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/el7onr/go-todo/config"
	"github.com/el7onr/go-todo/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	Config *config.Config
	Server *echo.Echo
}

func StartServer(c *config.Config) {

	r := echo.New()

	cors := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS", "HEADER"},
	}

	r.Use(middleware.CORSWithConfig(cors))
	r.Use(middleware.Logger())

	a := &Handler{
		Config: c,
		Server: r,
	}

	a.register(r)

	if err := r.Start("0.0.0.0:8080"); err != nil {
		log.Fatalf("error while starting server %v \n", err)
	}
}

func (h *Handler) register(e *echo.Echo) {
	e.POST("/todos", h.CreateToDo)
	e.GET("/todos", h.ListToDo)
	e.GET("/todos/:id", h.GetToDo)
	e.DELETE("/todos/:id", h.DeleteToDo)
	e.PUT("/todos/:id", h.UpdateToDo)
}

func (h *Handler) CreateToDo(c echo.Context) error {
	m := &model.ToDo{}

	err := c.Bind(m)
	if err != nil {
		c.JSON(http.StatusBadRequest, generateErrorResponse("invalid json", err))
		return err
	}

	r, err := h.Config.DB.CreateToDo(m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorResponse("error while writing data", err))
		return err
	}

	c.JSON(http.StatusOK, r)

	return nil
}

func (h *Handler) ListToDo(c echo.Context) error {
	r := h.Config.DB.ListToDo()

	c.JSON(http.StatusOK, r)

	return nil
}

func (h *Handler) GetToDo(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, generateErrorResponse("invalid parameter", err))
		return err
	}

	r, err := h.Config.DB.ReadToDo(id)
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

func (h *Handler) DeleteToDo(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, generateErrorResponse("invalid parameter", err))
		return err
	}

	err = h.Config.DB.DeleteToDo(id)
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

func (h *Handler) UpdateToDo(c echo.Context) error {
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

	r, err := h.Config.DB.UpdateToDo(m)
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
