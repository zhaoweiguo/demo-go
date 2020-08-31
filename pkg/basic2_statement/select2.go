package main

import (
	"fmt"
	"time"
)

var ch1 chan int
var ch2 chan int
var chs = []chan int{}
var numbers = []int{}

/*
所有channel表达式都会被求值、所有被发送的表达式都会被求值。求值顺序：自上而下、从左到右.
又因为没有消费消息，阻塞住，所以select语句中的fmt语句没有执行
*/
func main() {

	ch1 = make(chan int)
	ch2 = make(chan int)
	chs = []chan int{ch1, ch2}
	numbers = []int{1, 2, 3, 4, 5}
	for true {
		select {
		case getChan(0) <- getNumber(2):
			fmt.Println("1th case is selected.")
		case getChan(1) <- getNumber(3):
			fmt.Println("2th case is selected.")
		default:
			fmt.Println("default!.")
		}
		time.Sleep(time.Second)

	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)

	return numbers[i]
}
func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n", i)

	return chs[i]
}
