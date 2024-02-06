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
	var requested string = r.FormValue("q") // Get queried word from /define?q= 

	log.Printf("Got definition request for \"%s\"", requested)
	dictResponse := dictionary.Define(requested) // Look the word up, get back a Phrase struct

	// Prepare the template with corresponding struct fields and dictionary result
	data := struct {
		Query      string
		Definition string 
		Usage      string
	}{
		strings.Join(dictResponse.Names, "; "),
		dictResponse.Definition,
		dictResponse.Usage,
	}

	t, err := template.ParseFiles("web/template/definition.html")
	if err != nil {
		log.Fatal("Error while parsing definition template: ", err)
	}
	// Send the processed template to ResponseWriter
	t.Execute(w, data)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Get random phrase
	dictRandom := dictionary.Random()

	// Prepare the template 
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
	// Send the processed template to ResponseWriter
	t.Execute(w, data)
}

func main() {
	// Get port from runtime arguments
	var port string = ":" + os.Args[1]

	staticServer := http.FileServer(http.Dir("web/static"))
	http.Handle("/static", staticServer)

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/define", defineHandler)

	log.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
