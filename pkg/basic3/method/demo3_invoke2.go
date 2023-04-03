package main

import (
	"fmt"
)

func main() {
	demo1_invoke2()
	demo2_invoke2()
}

/*
总结论:
	1. 如果是值接收者，实体类型的值和指针都可以实现对应的接口；如果是指针接收者，那么只有类型的指针能够实现对应的接口
	Methods Receivers					Values
	(t T)methodName()				    T and *T
	(t *T)methodName()					   &T

	2.类型的值只能实现值接收者的接口；指向类型的指针，既可以实现值接收者的接口，也可以实现指针接收者的接口:
	Values			Methods Receivers
	T					(t T)methodName()
	&T				(t T)methodName() and (t *T)methodName()
*/

func demo1_invoke2() {
	fmt.Println("===demo1============================")
	var c cat_value
	//值作为参数传递
	invoke(c)
	invoke(&c) // 结论1: 实体类型以值接收者实现接口的时候，不管是实体类型的值，还是实体类型值的指针，都实现了该接口
}

func demo2_invoke2() {
	fmt.Println("===demo2_user============================")
	var c cat_point

	//invoke(c)		// 报错: Cannot use 'c' (type cat_point) as type animal2
	invoke(&c) // 结论2: 实体类型以指针接收者实现接口的时候，只有指向这个类型的指针才被认为实现了该接口
}

//需要一个animal接口作为参数
func invoke(a animal2) {
	a.printInfo()
}

type animal2 interface {
	printInfo()
}

type cat_value int

//值接收者实现animal接口
func (c cat_value) printInfo() {
	fmt.Println("a cat")
}

type cat_point int

//指针接收者实现animal接口
func (c *cat_point) printInfo() {
	fmt.Println("a cat")
}
