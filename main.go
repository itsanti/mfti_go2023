package main

import (
	"net/http"
	"time"
)

type RetryableHTTPclient struct {
	Retries int
	Timeout time.Duration
	Debug   bool
}

func (r *RetryableHTTPclient) Get(url string) *http.Response {
	// implement me
}
