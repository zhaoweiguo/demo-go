package main

import (
	"fmt"
	"time"
)

/*
给channel增加缓冲区
*/
func main() {
	ch := make(chan int, 1)
	ch <- 1
	go func() {
		v := <-ch
		fmt.Println(v)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("2")
}
