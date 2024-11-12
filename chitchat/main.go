package main

import (
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r)

		if err != nil {
			generateHtml(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHtml(w, threads, "layout", "public.navbar", "index")
		}
	}

}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	files := http.FileServer(http.Dir("/public"))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
