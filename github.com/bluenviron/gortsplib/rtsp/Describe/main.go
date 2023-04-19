package main

import (
	"log"

	"github.com/bluenviron/gortsplib/v3"
	"github.com/bluenviron/gortsplib/v3/pkg/url"
)

func main() {

	c := gortsplib.Client{}

	// parse URL
	u, err := url.Parse("rtsp://admin:admin123@192.168.124.2:554/live/ch0")
	if err != nil {
		panic(err)
	}

	// connect to the server
	err = c.Start(u.Scheme, u.Host)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// find published medias
	medias, baseURL, resp, err := c.Describe(u)
	if err != nil {
		panic(err)
	}
	log.Println(baseURL, resp)
	for _, media := range medias {
		log.Println(media)
	}
}
