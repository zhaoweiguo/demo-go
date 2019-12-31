package main

import (
	"fmt"
)

func main() {
	demo_slice1()
	demo_slice2()
}

func demo_slice1() {
	pcs := make([]uintptr, 2)
	fmt.Println(len(pcs)) // 2
}

func demo_slice2() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &slice)
	modify_slice(slice)
	fmt.Println(slice)

	/*
		输出:
		   0xc0000c6000
		   0xc00009a020
		   [1 10 3 4 5]

		查看:
			两个切片的地址不一样，所以可以确认切片在函数间传递是复制的。
			而我们修改一个索引的值后，发现原切片的值也被修改了，说明它们共用一个底层数组

		说明:
			在函数间传递切片非常高效，而且不需要传递指针和处理复杂的语法，只需要复制切片，然后根据自己的业务修改，最后传递回一个新的切片副本即可
			这也是为什么函数间传递参数，使用切片，而不是数组的原因
	*/
}

func modify_slice(slice []int) {
	fmt.Printf("%p\n", &slice)
	slice[1] = 10
}
