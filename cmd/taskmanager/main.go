package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1saswata/go-mentorship/internal/handlers"
	"github.com/1saswata/go-mentorship/internal/middleware"
	"github.com/1saswata/go-mentorship/internal/store"
)

func main() {
	db := InitDB()
	defer db.Close()
	mux := http.NewServeMux()
	wrappedMux := middleware.LoggingMiddleware(mux)
	tasks := handlers.TaskServer{Store: store.NewTaskStore(db)}
	mux.HandleFunc("/health", handlers.HealthCheckHandler)
	mux.HandleFunc("GET /tasks", tasks.ListTaskHandler)
	mux.HandleFunc("POST /tasks", tasks.CreateTaskHandler)
	mux.HandleFunc("PUT /tasks/{id}", tasks.UpdateTaskHandler)
	mux.HandleFunc("DELETE /tasks/{id}", tasks.DeleteTaskHandler)
	newServer := http.Server{Addr: ":8080", Handler: wrappedMux}
	c := make(chan os.Signal, 1)
	go func() {
		if err := newServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("HTTP Server error : ", err)
		}
	}()
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := newServer.Shutdown(ctx)
	if err != nil {
		log.Fatal("Error shutting down the server: ", err)
	}
	log.Print("Server is closed.")
}
