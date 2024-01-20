package main

import (
	"html/template"
	"net/http"
	"os"
)

func defineHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	data := struct {
		Query      string
		Definition string
	}{
		r.Form["q"][0],
		"test",
	}

	t, _ := template.ParseFiles("web/template/definition.html")
	t.Execute(w, data)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	data := struct {
		DailyPhrase     string
		DailyDefinition string
	}{
		"Daily test",
		"Daily test definition",
	}

	t, _ := template.ParseFiles("web/template/index.html")
	t.Execute(w, data)
}

func main() {
	http.Handle("/style.css", http.FileServer(http.Dir("web/static")))

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/define", defineHandler)
	http.ListenAndServe(":"+os.Args[1], nil)
}
