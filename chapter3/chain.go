package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func helloChain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HEllo with chain")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("FUNCTION Calleds" + name)
		h(w, r)
	}
}
