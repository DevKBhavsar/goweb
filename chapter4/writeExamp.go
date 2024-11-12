package main

func WriteExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go Web </title></head>
	<body><h1> HEllo World </h1></body>
	</html>`
	w.Write([]byte(str))
}

func WriteHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

type Profile struct {
	name   string
	thread []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	prof := &Profile{
		name:   "Dev Bhavar",
		thread: []string{"One", "TWO"},
	}
	js, _ := json.Marshal(prof)
	w.Write(js)

}
