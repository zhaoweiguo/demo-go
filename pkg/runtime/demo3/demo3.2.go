package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var i2 int64 = 0

/*
问题版2:
	通过共享内存来通信的改良版
	这儿的i值会变，但有可能出现:for循环执行1000次，go线程才执行一次的情况
	原因是: go线程每执行一次,sleep 1秒
*/
func main() {
	runtime.GOMAXPROCS(2)
	var m sync.Mutex
	go func() {
		for {
			m.Lock()
			fmt.Println("i2 is", i2)
			m.Unlock()
			time.Sleep(time.Second)
		}
	}()

	for {
		m.Lock()
		i2 += 1
		fmt.Println("i2 is for:", i2)
		m.Unlock()
	}
}
