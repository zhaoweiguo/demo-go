package main

import(
	"fmt"
)

func main() {
	as := []*A{}   // 这样的数据不会跟随
	b := &B{}
	b.init(as)

	as = append(as, &A{"a"})   // 这儿改变而改变
	as = append(as, &A{"b"})
	b.string()    // 这儿总是空的数组
}

type A struct {
	a string
}

type B struct {
	List []*A     // 
}

func (b B) init(as []*A) {
	b.List = as
}

func (b B) string() {
	fmt.Printf("b.List:%v\n", b.List)
}
