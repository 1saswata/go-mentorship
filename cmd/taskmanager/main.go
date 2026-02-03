package main

import (
	"io"
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK\n")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthCheckHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
