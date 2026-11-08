package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type Result struct {
	URL      string
	Up       bool
	Status   int
	Duration time.Duration
	Err      error
}

func checkURL(url string) Result {
	client := http.Client{Timeout: 5 * time.Second}
	start := time.Now()

	resp, err := client.Get(url)
	if err != nil {
		return Result{URL: url, Up: false, Err: err, Duration: time.Since(start)}
	}
	defer resp.Body.Close()

	return Result{
		URL:      url,
		Up:       resp.StatusCode >= 200 && resp.StatusCode < 400,
		Status:   resp.StatusCode,
		Duration: time.Since(start),
	}
}

func main() {
	urls := os.Args[1:]
	if len(urls) == 0 {
		fmt.Println("Usage: go run . <url1> <url2> ...")
		return
	}

	results := make(chan Result)
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			results <- checkURL(u)
		}(url)
	}

	// Close the channel once all goroutines finish
	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		if r.Up {
			fmt.Printf("UP    %s [%d] %v\n", r.URL, r.Status, r.Duration.Round(time.Millisecond))
		} else {
			fmt.Printf("DOWN  %s (%v)\n", r.URL, r.Err)
		}
	}
}
