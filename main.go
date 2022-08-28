package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexRoute)
	_ = http.ListenAndServe(":80", nil)
}

func indexRoute(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("template.html", "header.html"))
	_ = tmpl.ExecuteTemplate(w, "index", nil)
}
