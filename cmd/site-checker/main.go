package main

import (
	"log"
	"net/http"
	"sync"
)

func checkUrl(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get(url)
	if err == nil {
		log.Printf("%s - UP", url)
	} else {
		log.Printf("%s - DOWN", url)
	}
}

func main() {
	urls := []string{"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
		"http://doesntexist.com",
		"dadsadadw.dawds"}
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go checkUrl(url, &wg)
	}
	wg.Wait()
}
