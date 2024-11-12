package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Println(string(body))
	fmt.Fprintf(w, string(body))
}

func process(w http.ResponseWriter, r *http.Request) {
	file,_,err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintf(w, string(data))
		}
	}

}
