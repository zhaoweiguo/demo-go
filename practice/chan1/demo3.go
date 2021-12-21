package main

import (
	"fmt"
	"time"
)

func f(i int, input <-chan int, output chan<- int) {
	for {
		<-input
		fmt.Println(i)
		time.Sleep(time.Second)
		output <- 1
	}
}
func main() {
	c := [4]chan int{}
	for i := range []int{1, 2, 3, 4} {
		c[i] = make(chan int)
	}
	go f(1, c[3], c[0])
	go f(2, c[0], c[1])
	go f(3, c[1], c[2])
	go f(4, c[2], c[3])
	c[3] <- 1
	select {}
}
