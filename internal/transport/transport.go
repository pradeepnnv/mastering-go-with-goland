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
		b, err := json.Marshal(todoSvc.GetAll())
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

		todoSvc.Add(t.Item)
		//svc.Add(t)
		writer.WriteHeader(http.StatusCreated)
		return
	})
	return &Server{
		mux: mux,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}
