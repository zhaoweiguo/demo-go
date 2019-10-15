package main

import (
	"fmt"
	"time"
)

/*
如果有一个或多个IO操作可以完成，则Go运行时系统会随机的选择一个执行，
否则的话，如果有default分支，则执行default分支语句，
如果连default都没有，则select语句会一直阻塞，直到至少有一个IO操作可以进行
*/
func main() {
	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		var num int64
		num = 4 // 2. 大于3时, 会执行ch1或ch2这部分流程
		//num = 2   // 3. 小于3时, 会执行c这部分流程
		time.Sleep(time.Duration(num) * time.Second)
		close(c)
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 3
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 5
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c: // 3. num小于3,执行这儿
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	case <-ch1: // 2. 执行这儿，或下面ch2
		fmt.Printf("ch1 case...")
	case <-ch2:
		fmt.Printf("ch2 case...")
		//default:    // 1. 默认直接走这儿（如果上面都没有成功，则进入default处理流程）
		//	fmt.Printf("default go...")
	}
}
