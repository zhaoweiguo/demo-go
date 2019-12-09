package main

import (
	"fmt"
	"time"
)

func produce2(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}
func consumer2(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}
func main() {
	ch := make(chan int, 10)
	go produce2(ch)
	go consumer2(ch)
	time.Sleep(1 * time.Second)
}
