package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func helloRouter(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "HELLO %s", p.ByName("name"))
}
