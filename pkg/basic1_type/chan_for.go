package main

import (
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {
	do_select()
	//do_for()
}

/// select statement
func do_select() {
	ch := make(chan int, 10)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	go do_select2(ch)
	ch <- 100
	time.Sleep(time.Second * 3)
	ch <- 101
}

func do_select2(c chan int) {
	for true {
		do_select3(c)
	}
}

func do_select3(c chan int) {
	select {
	case v := <-c:
		log.Println(v)
	case <-time.After(1 * time.Second): // 因为有default,所以永远不会执行到这儿
		log.Println("time out")
		break
	default:
		log.Println("default")
		time.Sleep(time.Millisecond * 100)
	}
}

/// for statement
func do_for() {
	ch := make(chan int, 100)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	go do_for2(ch)
	for i := 10; i < 15; i++ {
		ch <- i
	}
	time.Sleep(time.Second)
}

func do_for2(ch <-chan int) {
	for data := range ch {
		log.Println(data)
	}
	log.Println("=======")
}
