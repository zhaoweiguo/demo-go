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
	//c := make(chan int64, 3)    // 加上缓存,可以看下与不加缓存有啥不同
	go func() {
		for {
			fmt.Println("i31 is", <-c)
			time.Sleep(time.Second)
		}
	}()

	for {
		i3 += 1
		fmt.Println("i32 is", i3)
		c <- i3
	}
}
