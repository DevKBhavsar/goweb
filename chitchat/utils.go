package main

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err := errors.New("INvalid Session")
		}
	}
	return
}
