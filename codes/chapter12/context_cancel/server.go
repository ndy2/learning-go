package main

import (
	"net/http"
	"net/http/httptest"
	"time"
)

// This server simulates a slow response by sleeping for 2 seconds before writing "Slow response" to the response writer.
func slowServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.Write([]byte("Slow response"))
	}))
	return s
}

// This server responds to HTTP requests based on the query parameter `error`.
func fastServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("error") == "true" {
			w.Write([]byte("error"))
			return
		}
		w.Write([]byte("ok"))
	}))
	return s
}
