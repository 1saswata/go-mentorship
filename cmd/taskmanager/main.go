package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK\n")
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var t Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func ListTaskHandler(w http.ResponseWriter, r *http.Request) {
	tasks := []Task{
		{1, "Feed the cat", "Complete"},
		{2, "Pet the cat", "Ongoing"},
		{5, "Let the cat sleep", "Incomplete"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthCheckHandler)
	mux.HandleFunc("GET /tasks", ListTaskHandler)
	mux.HandleFunc("POST /tasks", CreateTaskHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
