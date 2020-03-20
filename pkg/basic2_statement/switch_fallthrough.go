package main

import (
	"fmt"
)

func number() int {
	num := 15 * 5
	return num
}

/*
1. switch 与 case 中的表达式不必是常量，他们也可以在运行时被求值
2. fallthrough 必须是 case 语句块中的最后一条语句
*/
func main() {

	switch num := number(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num) // 打印
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num) // 打印
	}

}
