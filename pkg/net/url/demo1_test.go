package main

import (
	"fmt"
	"log"
	"net/url"
	"testing"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func TestA(t *testing.T) {
	u, err := url.Parse("http://www.baidu.com/search?q=dotnet")
	if err != nil {
		panic(err)
	}

	fmt.Printf("url.string(): %s\n", u.String())
}

func TestB(t *testing.T) {
	log.Println(url.ParseQuery("中国"))
	log.Println(url.PathEscape("aaa中国 北京bbb"))
}
