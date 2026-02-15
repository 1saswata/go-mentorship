package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func checkUrl(url string, wg *sync.WaitGroup, results chan string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("%s - DOWN (Network Error)", url)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		results <- fmt.Sprintf("%s - UP", url)
	} else {
		results <- fmt.Sprintf("%s - DOWN (Status: %d)", url, resp.StatusCode)
	}
}

func main() {
	urls := []string{"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
		"http://doesntexist.com",
		"dadsadadw.dawds"}
	var wg sync.WaitGroup
	results := make(chan string)
	for _, url := range urls {
		wg.Add(1)
		go checkUrl(url, &wg, results)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range results {
		log.Print(result)
	}
}
