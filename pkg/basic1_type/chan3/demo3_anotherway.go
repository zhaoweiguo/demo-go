package main

import "fmt"

/*
把ch<-1这一行代码放到子线程代码的后面:

*/
func main() {
	ch := make(chan int)

	go func() {
		v := <-ch
		fmt.Println(v)
	}()
	ch <- 1
	fmt.Println("2")
}
