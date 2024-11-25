package main

import (
	"log"
	"my-first-api/internal/db"
	"my-first-api/internal/todo"
	"my-first-api/internal/transport"
)

func main() {
	d, err := db.New("postgres", "mysecretpassword", "localhost", 5432, "postgres")
	if err != nil {
		log.Fatal(err)
	}
	svc := todo.NewService(d)
	server := transport.NewServer(svc)
	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
