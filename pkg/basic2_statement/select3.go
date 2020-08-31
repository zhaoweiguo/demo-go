package main

import (
	"fmt"
	"time"
)

/*
猜测:
select中的语句一个执行完之后，另一个不会执行。
所以同一个channel，但select只消费producer方法传入的数据,而不会消费自身select语句传入的数据

奇怪的是为啥consumer也不消费producer方法的数据呢？
*/

var (
	ch     chan int
	stopCh chan bool
)

// 2.1 作为生产者往ch生产数据
func producer(ch chan int, stopCh chan bool) {
	var i int
	i = 10
	for j := 0; j < 10; j++ {
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- true
}

// 2.2 作为消费者消费ch中的数据
func consumer(ch chan int) {
	for true {
		fmt.Println("~~~~>", <-ch)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	ch = make(chan int)
	c := 0
	stopCh = make(chan bool)

	go producer(ch, stopCh)
	go consumer(ch)

	for {
		select {
		// 1.1 作为消费者消费ch的数据
		case c = <-ch:
			fmt.Println("aaaa:", c)
		case s := <-ch:
			fmt.Println("bbbb:", s)
			// 1.2 作为生产者往ch生产数据
		case ch <- 2:
			fmt.Println("ccccc")
		case ch <- 1:
			fmt.Println("ccccc")
		case _ = <-stopCh:
			goto end
		}
	}
end:
}
