package main

import (
	"fmt"
)

func main() {
	demo1()
	demo2()
}

func demo1() {
	fmt.Println("===demo1============================")
	p := person{name: "张三"}
	p.modify() //值接收者，修改无效
	fmt.Println(p.String())
}

func demo2() {
	fmt.Println("===demo2============================")
	p := person2{name: "张三"}
	(&p).modify() //指针接收者，修改有效
	fmt.Println((&p).String())
}

type person struct {
	name string
}

func (p person) String() string {
	return "the person name is " + p.name
}

func (p person) modify() {
	p.name = "李四"
}

type person2 struct {
	name string
}

func (p *person2) String() string {
	return "the person name is " + p.name
}

func (p *person2) modify() {
	p.name = "李四"
}
