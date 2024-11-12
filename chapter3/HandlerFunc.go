package main

import (
	"fmt"
	"net/http"
)

//Working with Hander function

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "THIS IS HELLO HANDLE Function")
}
