package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
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
	if id == -1 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	json.NewEncoder(w).Encode(t)
}

func (ts *taskServer) ListTaskHandler(w http.ResponseWriter, r *http.Request) {
	tasks := ts.store.GetAllTasks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (ts *taskServer) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var t Task
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = ts.store.UpdateTaskStatus(id, t.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (ts *taskServer) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = ts.store.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	db := InitDB()
	defer db.Close()
	mux := http.NewServeMux()
	wrappedMux := LoggingMiddleware(mux)
	tasks := taskServer{store: NewTaskStore(db)}
	mux.HandleFunc("/health", HealthCheckHandler)
	mux.HandleFunc("GET /tasks", tasks.ListTaskHandler)
	mux.HandleFunc("POST /tasks", tasks.CreateTaskHandler)
	mux.HandleFunc("PUT /tasks/{id}", tasks.UpdateTaskHandler)
	mux.HandleFunc("DELETE /tasks/{id}", tasks.DeleteTaskHandler)
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
