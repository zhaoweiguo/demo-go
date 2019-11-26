package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("hello world")
	log.Fatal(http.ListenAndServe(":8084", nil))
}

type Index struct {
	Title string
	Body  string
}

// indexHandler is an HTTP handler that serves the index page.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	var indexTemplate = template.Must(template.ParseFiles("./template/index.tmpl"))
	log.Println("index Handler...", indexTemplate.Tree)
	data := &Index{
		Title: "Image gallery",
		Body:  "Welcome to the image gallery.",
	}

	if err := indexTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}
