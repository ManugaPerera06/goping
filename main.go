package main

import (
	"fmt"
	"net/http"
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
	urls := []string{
		"https://go.dev",
		"https://github.com",
		"https://google.com",
		"epicgames.com",
		"https://www.formula1.com/",
		"https://this-site-does-not-exist-12345.com",
	}

	for _, url := range urls {
		checkURL(url)
	}
}
