package main

import (
	"log"

	"github.com/bluenviron/gortsplib/v3"
	"github.com/bluenviron/gortsplib/v3/pkg/url"
	"github.com/zhaoweiguo/demo-go/github.com/bluenviron/gortsplib/rtsp/utils"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

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
	resp, err := c.Options(u)
	if err != nil {
		panic(err)
	}
	log.Println(resp)
}
