package main

import (
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// 锁的使用
var mutex = &sync.Mutex{} // or var mutex sync.Mutex
var totalTickets int32 = 10

func main() {
	runtime.GOMAXPROCS(4) //一个goroutine没被阻塞,别的goroutine就不会被执行。如想要真正并发需要加上这句
	rand.Seed(time.Now().Unix())

	for i := 0; i < 5; i++ {
		go sell_tickets(i)
	}

}

func sell_tickets(i int) {

}
