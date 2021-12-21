package main

import (
	"fmt"
)

func main() {
	fmt.Println("main start")
	defer run("1")
	noerror()
	defer run("2")
	defer func() {
		a := recover()
		fmt.Printf("inner func: %v\n", a)
	}()
	defer run("3")
	error()
	defer run("4")
}

func noerror() {
	fmt.Println("noerror start")
	defer run("noerror end with defer")
}

func error() {
	fmt.Println("error")
	defer run("error end with defer")
	panic("error")
}

func run(str string) {
	fmt.Printf("%v\n", str)
}
