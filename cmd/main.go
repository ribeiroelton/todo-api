package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/el7onr/go-todo/storage"
	"github.com/el7onr/go-todo/web/api"
	"github.com/labstack/echo/v4"
)

func main() {

	//Dependencies
	db := storage.NewDatabase()
	echo := echo.New()
	echo.HideBanner = true

	//Handlers
	api.NewApiHandler(echo, db)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("starting server")
		if err := echo.Start("0.0.0.0:8000"); err != nil {
			log.Fatalf("error while starting server, details %v \n", err)
		}
	}()

	<-ch

	log.Println("stopping server")
	if err := echo.Close(); err != nil {
		log.Fatalf("error while starting server, details %v \n", err)
	}

}
