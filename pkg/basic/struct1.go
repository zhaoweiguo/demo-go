/*
1. 结构体类普通方法
*/
package main

import (
	"fmt"
)

type A struct {
	A1 int
}

type B struct {
	B1 *A
	B2 string
}

func (b *B) Change() {
	b.B1 = &A{3}
	b.B2 = "c"
	fmt.Println("change")
	fmt.Printf("(%v,%v)\n", b.B1, b.B2) // 3
	change2(b)
	change3(b.B1)
}

func change3(a *A) {
	fmt.Println("local change3")
	a.A1 = 5
}
func change2(b *B) {
	b.B1 = &A{6}
	b.B2 = "e"
	fmt.Println("local change2")
	fmt.Printf("(%v,%v)\n", b.B1, b.B2) // 6
}

func (b *B) Check() {
	fmt.Println("check")
	fmt.Printf("(%v,%v)\n", b.B1, b.B2) // 5
}

func main() {
	a := &A{1}
	fmt.Printf("a:[%T]%v\n", a, a)   // a:[*main.A]&{1}
	fmt.Printf("a:[%T]%v\n", *a, *a) // a:[main.A]{1}
	b := &B{a, "a"}
	fmt.Printf("(%v,%v)\n", b.B1, b.B2) // 3: (&{1},a)
	b.Change()
	fmt.Printf("(%v,%v)\n", b.B1, b.B2) // 3
	b.Check()
	a.A1 = 2
	fmt.Printf("(%v,%v)\n", b.B1, b.B2) // 5
	change(b.B1)
	fmt.Printf("(%v,%v)\n", b.B1, b.B2) // 4
}

func change(a *A) {
	fmt.Println("local change")
	a.A1 = 4
}
