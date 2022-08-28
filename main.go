package main

import (
	"html/template"
	"net/http"
)

port := os.Getenv("PORT")

if port == "" {
		log.Fatal("$PORT must be set")
	}

func main() {
	http.HandleFunc("/", indexRoute)
	_ = http.ListenAndServe(":" + port, nil)
}

func indexRoute(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("template.html", "header.html"))
	_ = tmpl.ExecuteTemplate(w, "index", nil)
}
