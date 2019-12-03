/*
1. 结构体指针方法
2.
*/
// +build ignore

package main

import (
	"fmt"
)

type A2 struct {
	a1 int
}

type B2 struct {
	b1 *A2
	b2 int
}

func (b B2) Change() {
	b.b1 = &A2{3}
	fmt.Println("change:")
	fmt.Println(b.b1) // 3
}

func (b B2) Check() {
	fmt.Println("check")
	fmt.Println(b.b1) // 1
}

func main() {
	a := &A2{1}
	b := B2{a, 10}
	fmt.Println(b.b1) // 1
	b.Change()
	fmt.Println(b.b1) // 1
	b.Check()
	a.a1 = 2
	fmt.Println(b.b1) // 2
	change22(b.b1)
	fmt.Println(b.b1) // 4
}

func change22(a *A2) {
	a.a1 = 4
}
