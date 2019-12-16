package main

import (
	"fmt"
)

func main() {
	fmt.Println("before")
	panic("error")
	fmt.Println("after")
}
