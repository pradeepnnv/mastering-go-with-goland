package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type TodoItem struct {
	Id   int    `json:"id"`
	Item string `json:"item"`
}

var todos = make([]TodoItem, 0)

func main() {

	// persistent id to keep track of new todos being added to the list
	var pId int
	mux := http.NewServeMux()
	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(todos)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(b)
		if err != nil {
			log.Println()
		}
	})

	mux.HandleFunc("POST /todo", func(writer http.ResponseWriter, request *http.Request) {
		var t TodoItem
		err := json.NewDecoder(request.Body).Decode(&t)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		pId++
		t.Id = pId
		todos = append(todos, t)
		//svc.Add(t)
		writer.WriteHeader(http.StatusCreated)
		return
	})

	mux.HandleFunc("DELETE /todo", func(writer http.ResponseWriter, request *http.Request) {
		idParam := request.URL.Query().Get("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		// find the id in the todos list
		found := false
		i := 0
		for _, t := range todos {
			i++
			//log.Printf("param id is %d and todo id is %d and searching at %d", id, t.Id, i)
			if t.Id == id {
				found = true
				//log.Printf("found the todo to be deleted at %d whose id is %d", i, t.Id)
				break
			}
		}
		nTodos := make([]TodoItem, 0)
		nTodos = append(nTodos, todos[0:i-1]...)
		nTodos = append(nTodos, todos[i:]...)
		//log.Println("todos are ", todos, "new todos are ", nTodos)
		todos = nTodos
		if !found {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		writer.WriteHeader(http.StatusAccepted)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
