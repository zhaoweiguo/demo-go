package main

import (
	"fmt"
)

func main() {
	demo_struct1()
	demo_struct2()
	demo_struct3()
}

func demo_struct1() {
	fmt.Println("===demo_struct1============================")
	jim := person{10, "Jim"}
	fmt.Println(jim)
	modify_struct1(jim)
	fmt.Println(jim)
	/*
		jim2次的值是相同的
			传递的是值的副本
	*/
}

func demo_struct2() {
	fmt.Println("===demo_struct2============================")
	jim := person{10, "Jim"}
	fmt.Println(jim)
	modify_struct2(&jim)
	fmt.Println(jim)
	/*
		age的值已经被改变。
			传递的是结构体的指针
	*/
}

func demo_struct3() {
	fmt.Println("===demo_struct3============================")
	jim := person2{10, map[string]int{"a": 1}}
	fmt.Println(jim)
	modify_struct3(jim)
	fmt.Println(jim)
	/*
		如果结构体里有引用类型的值，比如map，
			那么我们即使传递的是结构体的值副本，如果修改这个map的话，原结构的对应的map值也会被修改
	*/
}

func modify_struct1(p person) {
	p.age = p.age + 10
}

func modify_struct2(p *person) {
	p.age = p.age + 10
}

func modify_struct3(p person2) {
	p.name["a"] = 2
}

type person struct {
	age  int
	name string
}
type person2 struct {
	age  int
	name map[string]int
}
