package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func checkUrl(ctx context.Context, url string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- fmt.Sprintf("%s - DOWN (Bad Request) %d", url, http.StatusBadRequest)
	}
	resp, err := http.DefaultClient.Do(req)
	select {
	case <-ctx.Done():
		results <- fmt.Sprintf("%s - DOWN (Status: %d)", url, http.StatusRequestTimeout)
		return
	default:
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

}

func main() {
	urls := []string{"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
		"http://doesntexist.com",
		"dadsadadw.dawds"}
	var wg sync.WaitGroup
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 900*time.Millisecond)
	defer cancel()
	results := make(chan string)
	for _, url := range urls {
		wg.Add(1)
		go checkUrl(ctx, url, &wg, results)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range results {
		log.Print(result)
	}
}
