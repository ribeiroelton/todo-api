package api

import (
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/el7onr/go-todo/config"
	"github.com/el7onr/go-todo/model"
	"github.com/el7onr/go-todo/storage"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	h        apiHandler
	ToDOJSON = `{
		"title": "title1",
		"description": "description1",
		"responsible": "responsible1",
		"when": "2014-05-16T08:28:06-00:00"
	}`
)

func setupEnv() {
	echoServer := echo.New()
	h = apiHandler{
		config: &config.Config{DB: nil},
		server: echoServer,
	}

}

func TestMain(m *testing.M) {
	setupEnv()
	os.Exit(m.Run())
}

func TestCreateToDo(t *testing.T) {
	db := storage.NewDatabase()
	h.config.DB = db

	req := httptest.NewRequest(echo.GET, "/todos", strings.NewReader(ToDOJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := h.server.NewContext(req, rec)

	assert.NoError(t, h.createToDo(c))

	t.Cleanup(func() {
		h.config.DB.DeleteToDo(0)
	})

}

func TestDeleteToDo(t *testing.T) {
	db := storage.NewDatabase()
	h.config.DB = db
	m := &model.ToDo{}

	h.config.DB.CreateToDo(m)

	req := httptest.NewRequest(echo.DELETE, "/todos", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := h.server.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("0")

	assert.NoError(t, h.deleteToDo(c))

	list := h.config.DB.ListToDo()
	assert.Len(t, list, 0)

	t.Cleanup(func() {
		h.config.DB.DeleteToDo(0)
	})
}
