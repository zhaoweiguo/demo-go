package main

import (
	"fmt"
)

func main() {
	demo_interface1()
	a := demo_interface2()
	b := a.(string)
	fmt.Println(b)
}

func demo_interface1() {

	var a interface{} = "77"
	b := a.(string)
	fmt.Printf("[%T]%v\n", b, b)
}

func demo_interface2() interface{} {
	return "abc"
}
