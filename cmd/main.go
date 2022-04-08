package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/el7onr/go-todo/config"
	"github.com/el7onr/go-todo/storage"
	"github.com/el7onr/go-todo/web/api"
)

func main() {
	log.Println("starting server")
	db := storage.NewDatabase()

	c := &config.Config{
		DB: db,
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	go api.StartServer(c)

	<-ch
	log.Println("stopping server")
}
