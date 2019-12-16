package main

import "fmt"

/*
报错:
fatal error: all goroutines are asleep - deadlock!
原因:
无缓冲的channel，然后给这个channel赋值了，程序就是在赋值完成后陷入了死锁
赋值后阻塞，等待值被取走，后面的步骤都不会被执行
解决:
见demo2与demo3
*/
func main() {
	ch := make(chan int)
	ch <- 1
	go func() {
		<-ch
		fmt.Println("1")
	}()
	fmt.Println("2")
}
