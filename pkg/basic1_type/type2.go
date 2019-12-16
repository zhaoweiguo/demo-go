package main

import (
	"fmt"
)

type abc interface{}
type d interface{}

func main() {
	a := "abc"
	b := (abc)(a)
	fmt.Printf("%v\n", b)

}
