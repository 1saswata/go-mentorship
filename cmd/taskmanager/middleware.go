package main

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		total := time.Since(start)
		log.Printf("%s %v %v", r.Method, r.URL, total)
	})
}
