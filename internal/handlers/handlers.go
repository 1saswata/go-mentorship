package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"sync"

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
	H     *Hub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Hub struct {
	clients   map[*websocket.Conn]bool
	broadcast chan []byte
	sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		msg := <-h.broadcast
		h.Lock()
		for c := range h.clients {
			c.WriteMessage(websocket.TextMessage, msg)
		}
		h.Unlock()
	}
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
	ts.H.Lock()
	ts.H.clients[conn] = true
	ts.H.Unlock()
	defer func() {
		_ = conn.Close()
		ts.H.Lock()
		delete(ts.H.clients, conn)
		ts.H.Unlock()
	}()
	go ts.H.Run()
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			slog.Error("Websocket error", "err", err)
			break
		}
		ts.H.broadcast <- p
	}
}
