package main

import (
	"log"

	"github.com/bluenviron/gortsplib/v3"
	"github.com/bluenviron/gortsplib/v3/pkg/url"
	"github.com/zhaoweiguo/demo-go/github.com/bluenviron/gortsplib/rtsp/utils"
)

func main() {

	//sender, _ := auth.NewSender([]string{"Basic realm=TP-Link IP-Camera"}, "admin", "admin123")
	c := gortsplib.Client{}
	// parse URL
	u, err := url.Parse(utils.Url)
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
		log.Println("==>", media.Type, media.Direction, media.Control)
		for _, f := range media.Formats {
			log.Println("\t\t-->", f.String())
		}
	}
}
