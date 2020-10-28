package main

import (
	"log"
)

func main() {
	demo_assign()
	demo_capable()
}

func demo_assign() {
	// 永远是符号<-进行读取或者写入，譬如v,ok := <-c是读取，而c <- value是写入
	c := make(chan int, 1)
	c <- 10  // 写入chan
	v := <-c // 从chan中读取
	log.Println(v)

	//判断chan是否关闭：
	c = make(chan int, 1)
	c <- 10
	v, ok := <-c // 读取，v=10，ok=true
	log.Println(v, ok)
	close(c)
	v, ok = <-c // 读取，v=0，ok=false
	log.Println(v, ok)
}

func demo_capable() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:3]
	log.Println(len(b), cap(b)) // 长度是3, capable是10
	//log.Println(b[5])		// 但5长度大于3超出index
}
