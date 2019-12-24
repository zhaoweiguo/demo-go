package main

import (
	"fmt"
	"runtime"
)

/*
这个函数的作用是让当前 goroutine 让出 CPU，当一个 goroutine 发生阻塞，
Go 会自动地把与该 goroutine 处于同一系统线程的其他 goroutine 转移到另一个系统线程上去，以使这些 goroutine 不阻塞
*/
func main() {
	proc1()
	proc2()
}

func proc1() {
	fmt.Println("==proc1=========================")
	runtime.GOMAXPROCS(1) //使用单核
	doit()
}

func proc2() {
	fmt.Println("==proc1=========================")
	runtime.GOMAXPROCS(4) //使用单核
	doit()
}

func doit() {
	exit := make(chan int)
	go func() {
		defer close(exit)
		go func() {
			fmt.Println("b")
		}()
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("a:", i)

		if i == 1 {
			runtime.Gosched() //切换任务
		}
	}
	<-exit
}
