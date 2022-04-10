package api_test

import (
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/el7onr/go-todo/config"
	"github.com/el7onr/go-todo/model"
	"github.com/el7onr/go-todo/storage"
	"github.com/el7onr/go-todo/web/api"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	handler  api.Handler
	ToDOJSON = `{
		"title": "title1",
		"description": "description1",
		"responsible": "responsible1",
		"when": "2014-05-16T08:28:06-00:00"
	}`
)

func setup() {
	echoServer := echo.New()
	handler = api.Handler{
		Config: &config.Config{DB: nil},
		Server: echoServer,
	}

}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestCreateToDo(t *testing.T) {
	db := storage.NewDatabase()
	handler.Config.DB = db

	req := httptest.NewRequest(echo.GET, "/todos", strings.NewReader(ToDOJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := handler.Server.NewContext(req, rec)

	assert.NoError(t, handler.CreateToDo(c))

	t.Cleanup(func() {
		handler.Config.DB.DeleteToDo(0)
	})

}

func TestDeleteToDo(t *testing.T) {
	db := storage.NewDatabase()
	handler.Config.DB = db
	m := &model.ToDo{}

	handler.Config.DB.CreateToDo(m)

	req := httptest.NewRequest(echo.DELETE, "/todos", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := handler.Server.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("0")

	assert.NoError(t, handler.DeleteToDo(c))

	list := handler.Config.DB.ListToDo()
	assert.Len(t, list, 0)

	t.Cleanup(func() {
		handler.Config.DB.DeleteToDo(0)
	})
}
