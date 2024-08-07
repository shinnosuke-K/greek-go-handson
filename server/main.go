package main

import (
	"encoding/json"
	"fmt"
	"greek-go-handson/repository"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// net/http

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("/hello")
		fmt.Fprint(w, "Hello World\n")
	})

	var repo repository.Repository
	repo = repository.NewInMemoryRepository()

	http.HandleFunc("/hello/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		type Body struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}
		var req Body
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if err := repo.Store(req.Key, req.Value); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	})
	http.HandleFunc("/hello/get", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value, err := repo.Get(key)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Fprint(w, value)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
