package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ribeiroelton/todo-api/internal/services"
	"github.com/ribeiroelton/todo-api/repository"
	"github.com/ribeiroelton/todo-api/web/api"
)

type server struct {
	echo *echo.Echo
}

func newServer() *server {
	return &server{
		echo: echo.New(),
	}
}

func (s *server) setupServer() {

	//Initialize Database
	db := repository.NewMemoryDB()

	//Initialize Service
	svc := services.NewToDoService(db)

	//Initialize Handler
	handler := api.NewToDoHandler(svc)

	//Configure Server
	s.echo.HideBanner = true

	cors := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS", "HEADER"},
	}

	s.echo.Use(middleware.CORSWithConfig(cors))
	s.echo.Use(middleware.Logger())

	s.echo.POST("/todos", handler.CreateToDo)
	s.echo.GET("/todos", handler.ListToDos)
	s.echo.GET("/todos/:id", handler.GetToDo)
	s.echo.DELETE("/todos/:id", handler.DeleteToDo)
	s.echo.PUT("/todos/:id", handler.UpdateToDo)
}

func (s *server) start(addr string) error {
	if err := s.echo.Start(addr); err != nil {
		return err
	}
	return nil
}

func (s *server) stop() error {
	if err := s.echo.Close(); err != nil {
		return err
	}
	return nil
}

func main() {

	server := newServer()
	server.setupServer()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("starting server")
		if err := server.start("0.0.0.0:8000"); err != nil {
			log.Fatalf("error while starting server, details %v \n", err)
		}
	}()

	<-ch

	log.Println("stopping server")
	if err := server.stop(); err != nil {
		log.Fatalf("error while starting server, details %v \n", err)
	}

}
