package main

import (
	"fmt"
)

func main() {
	demo1_byte()
	demo2_string()
	demo3_int()
	demo8_array_point()
	demo9_point_array()

	demo_assign()
}

func demo1_byte() {
	fmt.Println("===demo1_byte============================")
	// []byte, []uint8
	var val []byte = []byte(`"{\"c\":\"b\"}"`) // 这是切片
	for _, c := range val {
		fmt.Println(c)
	}
}

func demo2_string() {
	fmt.Println("===demo2_string============================")
	// []string
	subsCodes := []string{"aa", "vv", "dd", "ee", "gg"}
	for _, s := range subsCodes {
		fmt.Println(s)
	}

	// 数组新增元素
	cars := []string{"Gordon", "Simon", "Bland"}
	cars = append(cars, "Hope")
	fmt.Println(cars)

	car0 := cars[0]
	fmt.Println(car0)
	car1 := cars[1]
	fmt.Println(car1)
}

func demo3_int() {
	fmt.Println("===demo3_int============================")
	array := [5]int{1: 1, 3: 4}

	// 必须是长度一样，并且每个元素的类型也一样的数组，才是同样类型的数组
	var array1 [5]int = array //success
	//var array2 [4]int = array1 //error
	fmt.Println(array1, array)

}

// 数组指针
func demo8_array_point() {
	fmt.Println("===demo8_array_point============================")
	array := [5]int{1: 2, 3: 4}
	point := &array
	fmt.Println(point, array)
}

// 指针数组
func demo9_point_array() {
	//指针数组和数组本身差不多，只不过元素类型是指针
	fmt.Println("===demo9_point============================")
	array := [5]*int{1: new(int), 3: new(int)}
	fmt.Println(array[1], *array[1])
	*array[1] = 1
	fmt.Println(array[1], *array[1])

	//fmt.Println(array[0])	// ok
	//fmt.Println(*array[0])	// panic: runtime error: invalid memory address or nil pointer dereference
	array[0] = new(int)
	fmt.Println(array[0], *array[0]) // ok

}
