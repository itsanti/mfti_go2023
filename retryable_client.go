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
