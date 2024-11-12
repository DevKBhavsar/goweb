package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMesaage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("HELLO WORLD")
	c := http.Cookie{
		Name:  "Flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}

	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "no messge found")
		}
	} else {
		rc := http.Cookie{
			Name:    "Flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// http.HandleFunc("/h", Headers)
	// http.HandleFunc("/body", body)
	// http.HandleFunc("/process", process)
	// http.HandleFunc("/write", WriteExample)
	// http.HandleFunc("/writeHead", WriteHeaderExample)
	// http.HandleFunc("/json", jsonExample)
	// http.HandleFunc("/cookie", SetCookie)
	// http.HandleFunc("/get", GetCookie)

	http.HandleFunc("/setM", setMesaage)
	http.HandleFunc("/showM", showMessage)

	server.ListenAndServe()
}
