package main

import (
	"fmt"
	"runtime"
	"time"
)

var i int64 = 0

/*
问题版1:
会不断的输出0。全局变量i被两个goroutine同时读写，也就是我们所说的data race，导致了i的值是未定义的
*/
func main() {
	runtime.GOMAXPROCS(2) // 指定2个goroutine
	go func() {
		for {
			fmt.Println("i is", i)
			time.Sleep(time.Second)
		}
	}()

	for {
		i += 1
	}
}
