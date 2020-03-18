package main

import (
	"fmt"
)

// 字符串可视为char的数组
func main() {
	string3_1()
}

func string3_1() {
	var str string
	str = "abcdefghijk"
	a := str[0]
	fmt.Println(a)         // 97
	fmt.Println(a == 'a')  // true
	fmt.Println(97 == 'a') // true
}
