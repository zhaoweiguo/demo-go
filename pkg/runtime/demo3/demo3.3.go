package main

import (
	"fmt"
	"runtime"
	"time"
)

var i3 int64 = 0

/*
正式解决方法:
	通过通信来共享内存
*/
func main() {
	runtime.GOMAXPROCS(2)
	c := make(chan int64)
	go func() {
		for {
			fmt.Println("i3 is", <-c)
			time.Sleep(time.Second)
		}
	}()

	for {
		i3 += 1
		c <- i3
	}
}
