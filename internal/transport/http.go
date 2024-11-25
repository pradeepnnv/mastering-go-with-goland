package transport

import (
	"encoding/json"
	"log"
	"my-first-api/internal/todo"
	"net/http"
)

type TodoItem struct {
	Item string `json:"item"`
}

type Server struct {
	mux *http.ServeMux
}

func NewServer(todoSvc *todo.Service) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		todos, err := todoSvc.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		b, err := json.Marshal(todos)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(b)
		if err != nil {
			log.Fatal(err)
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

		err = todoSvc.Add(t.Item)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			_, err := writer.Write([]byte(err.Error()))
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		//svc.Add(t)
		writer.WriteHeader(http.StatusCreated)
		return
	})

	mux.HandleFunc("GET /search", func(writer http.ResponseWriter, request *http.Request) {
		query := request.URL.Query().Get("q")
		if query == "" {
			writer.WriteHeader(http.StatusBadRequest)
			log.Println("Query parameter is missing.")
			return
		}
		todos, err := todoSvc.Search(query)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if len(todos) == 0 {
			writer.WriteHeader(http.StatusNotFound)
			log.Println("No items found")
			return
		}
		b, err := json.Marshal(todos)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
		_, err = writer.Write(b)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

	})
	return &Server{
		mux: mux,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}
