package main

import (
	"fmt"
	"net/http"
)

func Headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Accept-Encoding"]
	fmt.Println(h)
	fmt.Fprintln(w, h)
}
