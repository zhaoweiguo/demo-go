package main

import (
	"fmt"
)

func main() {
	fmt.Println("main start")

	for i := 0; i < 5; i++ {
		fmt.Printf("i start:%v\n", i)
		defer logs("i end:%v\n", i)
	}
	fmt.Println("main end")
}

func logs(str string, i int) {
	fmt.Printf(str, i)
}
