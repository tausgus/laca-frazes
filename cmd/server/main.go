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

	tpl, _ := os.ReadFile("web/template/definition.html")
	t, _ := template.New("definition").Parse(string(tpl))
	t.Execute(w, data)
}

func main() {
	root := http.FileServer(http.Dir("web/static"))
	http.Handle("/", root)
	http.HandleFunc("/define", defineHandler)
	http.ListenAndServe(":"+os.Args[1], nil)
}
