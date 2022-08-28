package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", indexRoute)
	_ = http.ListenAndServe(":"+port, nil)
}

func indexRoute(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("template.html", "header.html"))
	_ = tmpl.ExecuteTemplate(w, "index", nil)
}
