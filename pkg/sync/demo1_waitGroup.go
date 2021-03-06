package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	demo1()
	demo2()
}

/*
说明1: sync.WaitGroup其实是一个计数的信号量
使用它的目的是要main函数等待两个goroutine执行完成后再结束，不然这两个goroutine还在运行的时候，程序就结束了，看不到想要的结果
说明2: 因为是并发执行，所以会打印一部分A，再打印一部分B，交替进行
*/
func demo1() {
	fmt.Println("===demo1============================")
	//runtime.GOMAXPROCS(1)	// 高级操作1: 强制指定一个逻辑处理器时和使用多个逻辑处理器有什么不同

	var wg sync.WaitGroup
	wg.Add(2) // 接收到2个wg.Done()后会结束
	go func() {
		defer wg.Done()
		doit("A")
	}()
	go func() {
		defer wg.Done()
		doit("B")
	}()
	wg.Wait()
}

func doit(tag string) {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(tag, ":[", i, j, "]->")
			//time.Sleep(time.Millisecond*1)	// 高级操作2: 逻辑处理器为1时，即执行: runtime.GOMAXPROCS(1)
			//因为goroutine执行时间太短暂了，还没来得及切换到第2个goroutine，第1个goroutine就完成了，所以显示的好像是顺序执行一样
			//可通过sleep让操作执行时间长一些
		}
		fmt.Println(tag, "=======>")
	}
}

func demo2() {
	fmt.Println("===demo2============================")
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
