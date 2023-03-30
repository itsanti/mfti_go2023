package main

import (
	"fmt"
	"net/http"
	"time"
)

type RetryableHTTPclient struct {
	Retries int
	Timeout time.Duration
	Debug   bool
}

func (r *RetryableHTTPclient) Get(url string) *http.Response {
	var response *http.Response
	var err error

	client := http.Client{
		Timeout: r.Timeout,
	}

	for i := 0; i <= r.Retries; i++ {
		response, err = client.Get(url)
		if response != nil {
			break
		}
		if r.Debug {
			fmt.Printf("Timeout exceeded: %v; Retry: %v; HTTP error\n", err, i+1)
		}
	}
	return response
}

func main() {
	const timeout = 2 * time.Second
	const retries = 3
	const debug = true
	const url = "http://example.com:81"

	c := RetryableHTTPclient{Timeout: timeout, Retries: retries, Debug: debug}
	response := c.Get(url)
	if response == nil {
		fmt.Printf("Server %s failed to respond after %d retries\n", url, retries)
	}
}
