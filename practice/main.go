package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO DEV")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Calling log")
		h(w, r)
	}
}

func Protect(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling protect")
		h(w, r)
	}
}

func main() {
	Server := http.Server{
		Addr: "127.0.0.1:80",
	}
	http.HandleFunc("/hello", Protect(log(Protect(Hello))))
	Server.ListenAndServe()
}
