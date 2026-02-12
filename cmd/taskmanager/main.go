package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type taskServer struct {
	store *TaskStore
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK\n")
}

func (ts *taskServer) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var t Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := ts.store.CreateTask(t.Name, t.Status)
	t.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func (ts *taskServer) ListTaskHandler(w http.ResponseWriter, r *http.Request) {
	tasks := ts.store.GetAllTasks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	mux := http.NewServeMux()
	tasks := taskServer{store: NewTaskStore()}
	mux.HandleFunc("/health", HealthCheckHandler)
	mux.HandleFunc("GET /tasks", tasks.ListTaskHandler)
	mux.HandleFunc("POST /tasks", tasks.CreateTaskHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
