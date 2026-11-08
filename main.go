package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func checkURL(url string) {
	client := http.Client{Timeout: 5 * time.Second}

	start := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("DOWN  %s (%v)\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("UP    %s [%d] %v\n", url, resp.StatusCode, time.Since(start).Round(time.Millisecond))
}

func main() {
	urls := os.Args[1:]
	if len(urls) == 0 {
		fmt.Println("Usage: go run. <url1> <url2> ...")
		return
	}
	for _, url := range urls {
		checkURL(url)
	}
}
