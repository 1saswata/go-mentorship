package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "OK\n")
	}
	http.HandleFunc("/health", h1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
