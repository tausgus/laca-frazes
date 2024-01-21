package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/tausgus/laca-frazes/internal/dictionary"
)

func defineHandler(w http.ResponseWriter, r *http.Request) {
	var requested string = r.FormValue("q")
	log.Printf("Got definition request for \"%s\"", requested)

	data := struct {
		Query      string
		Definition string
	}{
		strings.ToLower(requested),
		dictionary.Define(requested),
	}

	t, err := template.ParseFiles("web/template/definition.html")
	if err != nil {
		log.Fatal("Error while parsing definition template: ", err)
	}
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

	t, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatal("Error while parsing index template: ", err)
	}
	t.Execute(w, data)
}

func main() {
	var port string = ":" + os.Args[1]

	http.Handle("/style.css", http.FileServer(http.Dir("web/static")))

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/define", defineHandler)

	log.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
