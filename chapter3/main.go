package main

import (
	"fmt"
	"net/http"
)

type Final struct {
}

func (f *Final) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HEELO DEV")
}

func main() {

	tempServer := http.Server{
		Addr: "127.0.0.1:8080",
	}

	tempServer.ListenAndServe()
}
