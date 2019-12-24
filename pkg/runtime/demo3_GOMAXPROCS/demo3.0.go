package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4) //一个goroutine没被阻塞,别的goroutine就不会被执行。如想要真正并发需要加上这句
	rand.Seed(time.Now().Unix())

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		go sell_tickets(i)
	}
	time.Sleep(1 * time.Millisecond)

}

func sell_tickets(i int) {
	fmt.Println("ticket:", i)
	return
}
