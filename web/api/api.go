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

type apiServer struct {
	config *config.Config
	server *echo.Echo
}

func StartServer(c *config.Config) {

	r := echo.New()

	cors := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS", "HEADER"},
	}

	r.Use(middleware.CORSWithConfig(cors))
	r.Use(middleware.Logger())

	a := &apiServer{
		config: c,
		server: r,
	}

	a.register(r)

	if err := r.Start("0.0.0.0:8080"); err != nil {
		log.Fatalf("error while starting server %v \n", err)
	}
}

func (a *apiServer) register(e *echo.Echo) {
	e.POST("/todos", a.createToDo)
	e.GET("/todos", a.listToDo)
	e.GET("/todos/:id", a.getToDo)
	e.DELETE("/todos/:id", a.deleteToDo)
	e.PUT("/todos/:id", a.updateToDo)
}

func (a *apiServer) createToDo(c echo.Context) error {
	m := &model.ToDo{}

	err := c.Bind(m)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid json")
		return err
	}

	r, err := a.config.DB.CreateToDo(m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error while writing data")
		return err
	}

	c.JSON(http.StatusOK, r)

	return nil
}

func (a *apiServer) listToDo(c echo.Context) error {
	r := a.config.DB.ListToDo()

	c.JSON(http.StatusOK, r)

	return nil
}

func (a *apiServer) getToDo(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid parameter")
		return err
	}

	r, err := a.config.DB.ReadToDo(id)
	if err != nil && err.Error() == "id not found" {
		c.NoContent(http.StatusNotFound)
		return err
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error while getting data")
		return err
	}

	c.JSON(http.StatusOK, r)

	return nil
}

func (a *apiServer) deleteToDo(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid parameter")
		return err
	}

	err = a.config.DB.DeleteToDo(id)
	if err != nil && err.Error() == "id not found" {
		c.NoContent(http.StatusNotFound)
		return err
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error while deleting data")
		return err
	}

	c.NoContent(http.StatusOK)

	return nil

}

func (a *apiServer) updateToDo(c echo.Context) error {
	m := &model.ToDo{}

	err := c.Bind(m)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid json")
		return err
	}

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid parameter")
		return err
	}

	m.Id = id

	r, err := a.config.DB.UpdateToDo(m)
	if err != nil && err.Error() == "id not found" {
		c.NoContent(http.StatusNotFound)
		return err
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, "error while updating data")
		return err
	}

	c.JSON(http.StatusAccepted, r)
	return nil

}
