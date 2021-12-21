package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main() {
	var g errgroup.Group
	var urls = []string{
		"http://www.163.com/",
		"http://www.baidu.com/",
		"http://www.google.com/",                // 增加一个不可请求的网址，会在后面打印出失败原因
		"http://www.fasldfjaffsadfkjsjfsk.com/", // 增加第2个不可用网址,后面只打印最后一个失败网址
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			fmt.Println(url)
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println("fail", err)
	}

}
