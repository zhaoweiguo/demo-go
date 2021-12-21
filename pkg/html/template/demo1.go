package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("start... http://127.0.0.1:8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}

type Index struct {
	Title string
	URL   string
}

// indexHandler is an HTTP handler that serves the index page.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	var indexTemplate = template.Must(template.ParseFiles("./html/template/template/index.tmpl"))
	log.Println("index Handler...", indexTemplate.Tree)
	data := &Index{
		Title: "Image gallery",
		URL:   "http://www.baidu.com",
	}

	if err := indexTemplate.ExecuteTemplate(w, "content", data); err != nil {
		log.Println(err)
	}
}
