package main

import (
	"fmt"
)

func main() {
	as := []*A4{} // 这样的数据不会跟随
	b := &B4{}
	b.init(as)

	as = append(as, &A4{"a"}) // 这儿改变而改变
	as = append(as, &A4{"b"})
	b.string() // 这儿总是空的数组
}

type A4 struct {
	a string
}

type B4 struct {
	List []*A4 //
}

func (b B4) init(as []*A4) {
	b.List = as
}

func (b B4) string() {
	fmt.Printf("b.List:%v\n", b.List)
}
