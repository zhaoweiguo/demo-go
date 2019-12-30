package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mailgun/scroll"
)

func handler(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
	return scroll.Response{
		"message": fmt.Sprintf("Resource ID: %v", params["resourceID"]),
	}, nil
}

func handleError(w http.ResponseWriter, r *http.Request) {
	scroll.ReplyError(w, scroll.NotFoundError{Description: "Object not found"})
}

func handleTest1(w http.ResponseWriter, r *http.Request, params map[string]string, body []byte) (interface{}, error) {
	return scroll.Response{
		"Status": "ok",
	}, nil
}

func main() {
	// create an app
	app, err := scroll.NewApp()
	if err != nil {
		log.Println(err)
	}

	app.SetNotFoundHandler(handleError)
	app.AddHandler(scroll.Spec{Paths: []string{"/test1"}, Methods: []string{"GET"}, HandlerWithBody: handleTest1})

	server := &http.Server{
		Addr:    "0.0.0.0:7101",
		Handler: app.GetHandler(),
	}

	server.ListenAndServe()
}
