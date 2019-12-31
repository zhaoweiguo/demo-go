package main

import (
	"fmt"
)

// 多态
func main() {
	var a animal

	var c cat
	a = c
	a.printInfo()
	//使用另外一个类型赋值
	var d dog
	a = d
	a.printInfo()
}

type animal interface {
	printInfo()
}

type cat int
type dog int

func (c cat) printInfo() {
	fmt.Println("a cat")
}

func (d dog) printInfo() {
	fmt.Println("a dog")
}
