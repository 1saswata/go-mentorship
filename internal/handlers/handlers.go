package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/1saswata/go-mentorship/internal/store"
	"github.com/gorilla/websocket"
)

type Store interface {
	CreateTask(string, string) int
	GetAllTasks() []store.Task
	UpdateTaskStatus(int, string) error
	DeleteTask(int) error
}

type TaskServer struct {
	Store Store
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "OK\n")
}

func (ts *TaskServer) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var t store.Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := ts.Store.CreateTask(t.Name, t.Status)
	t.ID = id
	w.Header().Set("Content-Type", "application/json")
	if id == -1 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(t)
	}
}

func (ts *TaskServer) ListTaskHandler(w http.ResponseWriter, r *http.Request) {
	tasks := ts.Store.GetAllTasks()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(tasks)
}

func (ts *TaskServer) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var t store.Task
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = ts.Store.UpdateTaskStatus(id, t.Status)
	if err != nil {
		if err == store.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (ts *TaskServer) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = ts.Store.DeleteTask(id)
	if err != nil {
		if err == store.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (ts *TaskServer) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for {
		mt, p, err := conn.ReadMessage()
		if err != nil {
			slog.Error("Websocket error", "err", err)
			break
		}
		err = conn.WriteMessage(mt, p)
		if err != nil {
			slog.Error("Websocket error", "err", err)
			break
		}
	}
	_ = conn.Close()
}
