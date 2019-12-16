package main

import (
	"fmt"
)

func main() {
	fmt.Println("before")
	defer func() {
		err := recover()
		fmt.Println(err) // error
	}()
	panic("error")
	fmt.Println("after")
}
