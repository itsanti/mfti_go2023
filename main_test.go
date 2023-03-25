package main

import (
	"testing"
	"time"
)

const timeout = 2 * time.Second
const retries = 3
const debug = true

func TestRetryableHTTP(t *testing.T) {
	c := RetryableHTTPclient{Timeout: timeout, Retries: retries, Debug: debug}
	start := time.Now()
	_ = c.Get("http://example.com:81")
	elapsed := time.Since(start)
	seconds := elapsed.Seconds()
	want := timeout.Seconds() * retries
	if int(seconds) != int(want) {
		t.Errorf("elapsed == %v sec; want %v", seconds, want)
	}
}

func TestRetryableHTTPok(t *testing.T) {
	c := RetryableHTTPclient{Timeout: timeout, Retries: retries, Debug: debug}
	start := time.Now()
	r := c.Get("http://example.com")
	elapsed := time.Since(start)
	seconds := elapsed.Seconds()
	want := timeout.Seconds() * retries
	if seconds >= want {
		t.Errorf("elapsed >= %v sec; want %v", want, seconds)
	}
	if r.StatusCode != 200 {
		t.Errorf("response: %s", r.Status)
	}
}
