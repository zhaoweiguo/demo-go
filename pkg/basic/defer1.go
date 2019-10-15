package main

import (
	"fmt"
)

func main() {
	fmt.Println("main start")
	defer run(1)
	noerror()
	defer run(2)
	defer func () {
		a := recover()
		fmt.Printf("a: %v\n", a)
	}()
	defer run(3)
	error()
	defer run(4)
}

func noerror() {
	fmt.Println("noerror")
	defer run(11)
}

func error() {
	fmt.Println("error")
	defer run(21)
	panic("error")
}

func run(i int) {
	fmt.Printf("%d\n", i)
}
