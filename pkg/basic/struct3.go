/*
1. 结构体指针方法
2.
*/

package main

import (
	"fmt"
)

type A3 struct {
	A1 int
	B1 bool
	C1 string
}

// 此种用法主要用于封装
type C3 struct {
	*A3
}

func main() {
	a := A3{}
	fmt.Printf("a:%v\n", a)

	c := &C3{&a}
	c.A1 = 1
	c.B1 = true
	fmt.Println(c, c.A3, c.A1, c.A3.A1)
}
