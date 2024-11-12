package main

import (
	"fmt"
	"net/http"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    " SOME VALUE 1",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    " WOW SESCOND COOKIE",
		HttpOnly: true,
	}

	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)

}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	c1, _ := r.Cookie("first_cookie")
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)

}
