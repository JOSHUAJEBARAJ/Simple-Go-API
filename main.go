package main

import (
	"github.com/JOSHUAJEBARAJ/Simple-Go-API/internal/todo"
	"github.com/JOSHUAJEBARAJ/Simple-Go-API/internal/transport"
)

func main() {

	// var todos = make([]TodoItem, 0)
	svc := todo.NewService()
	server := transport.NewServer(svc)
	err := server.Serve()
	if err != nil {
		panic(err)
	}

}
