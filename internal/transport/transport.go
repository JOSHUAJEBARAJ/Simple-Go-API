package transport

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JOSHUAJEBARAJ/Simple-Go-API/internal/todo"
)

// used to marshal and demarshal
type TodoItem struct {
	Item string `json:"item"`
}

type Server struct {
	mux *http.ServeMux
}

func NewServer(svc *todo.Service) *Server {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {

		todos, err := svc.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(todos)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// json.NewEncoder(w).Encode(todos)
		return

	})

	mux.HandleFunc("POST /todo", func(w http.ResponseWriter, r *http.Request) {
		var i TodoItem
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = svc.Add(i.Item)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return

	})

	// mux.HandleFunc("POST /search", func(w http.ResponseWriter, r *http.Request) {
	// 	query := r.URL.Query().Get("q")
	// 	if query == "" {
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}
	// 	output := svc.Search(query)

	// 	b, err := json.Marshal(output)
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Write(b)
	// 	return

	// })

	return &Server{
		mux: mux,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8090", s.mux)
}
