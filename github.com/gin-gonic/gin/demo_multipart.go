package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//对应client: net/http/demo_client_file_upload
func main() {
	port := ":8844"
	readTimeOut := 60
	writeTimeOut := 60

	router := gin.Default()
	router.Any("/multipart", handle)

	s := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    time.Duration(readTimeOut) * time.Second,
		WriteTimeout:   time.Duration(writeTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("s=", *s)
	s.ListenAndServe()
}

func handle(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		log.Println(err)
	}
	value := form.Value
	file := form.File
	for k, v := range value {
		log.Println(k)
		log.Println("-------")
		log.Println(v)
		for _, v4 := range v {
			log.Println(v4)
		}
	}
	log.Println("=============")
	for k2, v2 := range file {
		log.Println(k2)
		log.Println(v2)
		for _, v3 := range v2 {
			log.Println(v3.Header)
			log.Println(v3.Size)
			log.Println(v3.Filename)
			f, err := v3.Open()
			if err != nil {
				log.Println(err)
			}
			p := make([]byte, v3.Size)
			n, err := f.Read(p)
			if err != nil {
				log.Println(err)
			}
			log.Println(n)
			log.Println(string(p))
			defer f.Close()
		}
		log.Println("-------")
	}
}
