package array

import (
	"fmt"
	"testing"
)

// 传递数组
func TestInvoke1(t *testing.T) {
	// 在函数间传递变量时，总是以值的方式，如果变量是个数组，那么就会整个复制，并传递给函数，
	//如果数组非常大，比如长度100多万，那么这对内存是一个很大的开销。
	fmt.Println("===demo1============================")
	array := [5]int{1: 2, 3: 4}
	modify(array)
	fmt.Println(array)
}

// 传递数组指针
func TestInvoke2(t *testing.T) {
	fmt.Println("===demo2_user============================")
	array := [5]int{1: 2, 3: 4}
	modify2(&array)
	fmt.Println(array)
}

func modify(a [5]int) {
	a[1] = 3
	fmt.Println(a)
}
func modify2(a *[5]int) {
	a[1] = 3
	fmt.Println(*a)
}
