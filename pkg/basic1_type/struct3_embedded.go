/*
1. 结构体类型嵌入
2. 结构体指针方法
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

func (a3 A3) Func_struct1() {
	fmt.Println("[fun]...a3 func")
}

// 此种用法主要用于封装
type C3 struct {
	*A3 // 结构体类型嵌入
}

func main() {
	struct3_1()
	struct3_2()
}

// C3可以直接用A3的变量和方法
func struct3_1() {
	fmt.Println("=========================直接使用A3的方法&变量:")
	a := A3{}
	fmt.Printf("a:%v\n", a)

	c := &C3{&a}
	c.A1 = 1 // 直接使用A3的变量
	c.B1 = true
	c.Func_struct1() // 直接使用A3的方法
	fmt.Println(c.A3, c.A1, c.A3.A1)
}

// 此种用法主要用于封装
type CC3 struct {
	*A3 // 结构体类型嵌入
	A1  int
}

func (a3 CC3) Func_struct1() {
	fmt.Println("[fun]... cc3 func")
}

func struct3_2() {
	fmt.Println("=========================A3和CC3的方法&变量相同时使用CC3的:")
	a := A3{}
	fmt.Printf("a:%v\n", a)
	cc := &CC3{&a, 5}
	cc.Func_struct1()
	fmt.Println(cc.A3, cc.A1, cc.A3.A1)
}
