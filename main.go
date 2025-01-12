package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type TodoItem struct {
	Item string `json:"item"`
}

func main() {
	mux := http.NewServeMux()
	var todos = make([]TodoItem, 0)

	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {

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
		todos = append(todos, i)
		w.WriteHeader(http.StatusCreated)
		return

	})

	if err := http.ListenAndServe(":8090", mux); err != nil {
		panic(err)
	}

}
