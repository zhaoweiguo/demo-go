package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	var indexTemplate = template.Must(template.ParseFiles("./template/index.tmpl"))
	data := &Index2{
		Title: "Image gallery",
		URL:   "http://www.baidu.com",
	}

	// 注意: "content"就是模板中的{{define "content"}}
	if err := indexTemplate.ExecuteTemplate(os.Stdout, "content", data); err != nil {
		log.Println(err)
	}
}

type Index2 struct {
	Title string
	URL   string
}
