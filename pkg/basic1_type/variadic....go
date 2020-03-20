package main

import "fmt"

// 这个函数可以传入任意数量的整型参数
func sum1(nums ...int) {
	fmt.Print(nums)
}

func sum2(nums []int) {
	fmt.Println(nums)
}

func main() {

	// 支持可变长参数的函数调用方法和普通函数一样
	// 也支持只有一个参数的情况
	sum1(1, 2)
	sum1(1, 2, 3)

	// 如果你需要传入的参数在一个切片中，像下面一样
	// "func(slice...)"把切片打散传入
	nums := []int{1, 2, 3, 4}
	sum1(nums...)

	sum2(nums)
	//sum2(1, 2, 3 )

}
