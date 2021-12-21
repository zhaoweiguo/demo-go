package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	fmt.Println("start....")
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.baidu.com/",
		"http://www.qq.com/",
		"http://www.360.cn/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		fmt.Println(url)
		wg.Add(1)
		fmt.Println(url)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("%v\n", err)
			} else {
				fmt.Printf("[%v]%v\n", url, resp)
			}

		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
