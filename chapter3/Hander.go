package main

import (
	"fmt"
	"net/http"
)

// Working with Handler
type Hello struct {
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO DEV ON EVERY PAGE")
}

func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

type World struct {
}

func (s *World) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World DEV ON EVERY PAGE")
}
