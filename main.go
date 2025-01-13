package main

import (
	"log"

	"github.com/JOSHUAJEBARAJ/Simple-Go-API/internal/db"
	"github.com/JOSHUAJEBARAJ/Simple-Go-API/internal/todo"
	"github.com/JOSHUAJEBARAJ/Simple-Go-API/internal/transport"
)

func main() {

	// var todos = make([]TodoItem, 0)
	db, err := db.NewDB("admin", "root", "postgres", "localhost", 5432)
	if err != nil {
		log.Fatal(err)
	}
	svc := todo.NewService(db)
	server := transport.NewServer(svc)
	err = server.Serve()
	if err != nil {
		panic(err)
	}

}
