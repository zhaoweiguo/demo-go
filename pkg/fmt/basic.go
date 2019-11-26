package main

import "fmt"

/*
%T: 变量类型
%t: boole类型
%d: int类型
%s: string类型
%b: 二进制类型
%c: unicode码点值
%v: 使用默认格式的内置或者自定义类型的值, 或者是此类型的String()输出
*/
func main() {

	var a int
	var b string
	var c bool

	a = 1
	b = "abc"
	c = true

	fmt.Printf("%T, %T, %T\n", a, b, c)
	fmt.Printf("%d\n", a)
	fmt.Printf("%s\n", b)
	fmt.Printf("%t\n", c)

}
