package main

import (
	"net/http"
	"text/template"

	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func custom(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("t1.html").Funcs(funcMap)
	t, _ = t.ParseFiles("t1.html")
	t.Execute(w, time.Now())
}
func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html")
	// content:= ` I Asked <i>What's up</i>`
	t.Execute(w,(r.FormValue("comment")))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t2.html")
	t.Execute(w, nil)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	http.HandleFunc("/form", form)

	server.ListenAndServe()
}
