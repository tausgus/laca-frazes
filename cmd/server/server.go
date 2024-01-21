package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/tausgus/laca-frazes/internal/dictionary"
)

func defineHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Query      string
		Definition string
	}{
		strings.ToLower(r.FormValue("q")),
		dictionary.Define(r.FormValue("q")),
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
