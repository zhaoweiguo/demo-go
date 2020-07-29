package main

import (
	"fmt"
	"log"
	"time"
)

/*
说明: 多个reader共同一个channel时
每个reader轮流获取数据
*/

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {

	c := make(chan string, 1)
	go sub1(c)
	go sub2(c)
	for i := 0; i < 50; i++ {
		c <- fmt.Sprint("num:~p", i)
		time.Sleep(time.Second * 1)
	}
}

func sub1(c chan string) {
	for true {
		a := <-c
		log.Println(a)
		time.Sleep(time.Second * 1)
	}
}

func sub2(c chan string) {
	for true {
		a := <-c
		log.Println(a)
		time.Sleep(time.Second * 1)
	}
}
