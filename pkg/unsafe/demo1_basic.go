package main

import (
	"fmt"
	"unsafe"
)

func main() {
	demo1_sizeof()
	demo1_pointer()
}

func demo1_sizeof() {
	// basic
	fmt.Println("===demo1_sizeof============================")
	fmt.Println(unsafe.Sizeof(true))                   // 大小1
	fmt.Println(unsafe.Sizeof(int8(0)))                // 大小1
	fmt.Println(unsafe.Sizeof(int16(10)))              // 大小2
	fmt.Println(unsafe.Sizeof(int32(10000000)))        // 大小4
	fmt.Println(unsafe.Sizeof(int64(10000000000000)))  // 大小8
	fmt.Println(unsafe.Sizeof(int(10000000000000000))) // 大小8(因为本系统是64位)

	// struct
	// 因为有内存对齐存在，编译器使用了内存对齐，那么最后的大小结果就不一样
	var u1 user1
	var u2 user2
	var u3 user3
	var u4 user4
	var u5 user5
	var u6 user6
	fmt.Println("u1 size is ", unsafe.Sizeof(u1)) // 大小16
	fmt.Println("u2 size is ", unsafe.Sizeof(u2)) // 大小24
	fmt.Println("u3 size is ", unsafe.Sizeof(u3)) // 大小16
	fmt.Println("u4 size is ", unsafe.Sizeof(u4)) // 大小24
	fmt.Println("u5 size is ", unsafe.Sizeof(u5)) // 大小16
	fmt.Println("u6 size is ", unsafe.Sizeof(u6)) // 大小16
}

/*
为了安全的考虑，Go语言是不允许两个指针类型进行转换
两个不同的指针类型不能相互转换，比如*int不能转为*float64
unsafe.Pointer是一种特殊意义的指针，它可以包含任意类型的地址，有点类似于C语言里的void*指针，全能型的
*/
func demo1_pointer() {
	fmt.Println("===demo1_pointer============================")
	i := 10
	ip := &i
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp / 3
	fmt.Println(*fp, *ip) // *ip值好理解, 而*fp的值是很大或很小的数
	// (原因是float64与int占空间大小不同, 指针相同只说明初使相同,具体数据大小要看后面空间当时存放的数据)
	fmt.Println(fp, ip)                         // fp==ip
	fmt.Println("修改的fp,但i的值也会改变(因为共用了同一指针)", i) // 相当于做了 i/3 操作

	// 字符串相关
	u := new(user)
	fmt.Println(*u)
	pName := (*string)(unsafe.Pointer(u)) // user结构的指针与user下name字段指针相同
	*pName = "张三"

	// 内存偏移牵涉到的计算只能通过uintptr，所我们要先把user的指针地址转为uintptr，
	//然后我们再通过unsafe.Offsetof(u.age)获取需要偏移的值，进行地址运算(+)偏移即可
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age))) // user下age字段指针等于user指针+u.age的偏移量
	*pAge = 20

	fmt.Println(*u)
}

type user1 struct {
	b byte
	i int32
	j int64
}

type user2 struct {
	b byte
	j int64
	i int32
}

type user3 struct {
	i int32
	b byte
	j int64
}

type user4 struct {
	i int32
	j int64
	b byte
}

type user5 struct {
	j int64
	b byte
	i int32
}

type user6 struct {
	j int64
	i int32
	b byte
}

type user struct {
	name string
	age  int
}
