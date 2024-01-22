package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/tausgus/laca-frazes/internal/dictionary"
)

func defineHandler(w http.ResponseWriter, r *http.Request) {
	var requested string = r.FormValue("q")

	log.Printf("Got definition request for \"%s\"", requested)
	dictResponse := dictionary.Define(requested)

	data := struct {
		Query      string
		Definition string
		Usage      string
	}{
		dictResponse.Names[0],
		dictResponse.Definition,
		dictResponse.Usage,
	}

	t, err := template.ParseFiles("web/template/definition.html")
	if err != nil {
		log.Fatal("Error while parsing definition template: ", err)
	}
	t.Execute(w, data)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	dictRandom := dictionary.Random()

	data := struct {
		DailyPhrase     string
		DailyDefinition string
		PhraseCount     int
	}{
		dictRandom.Names[0],
		dictRandom.Definition,
		dictionary.Stats(),
	}

	t, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatal("Error while parsing index template: ", err)
	}
	t.Execute(w, data)
}

func main() {
	var port string = ":" + os.Args[1]

	staticServer := http.FileServer(http.Dir("web/static"))

	http.Handle("/style.css", staticServer)
	http.Handle("/favicon.ico", staticServer)

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/define", defineHandler)

	log.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
