/*
1. 结构体指针方法
2. 
*/

package main

import(
	"fmt"
)

type A3 struct {
	A1  int
	B1  bool
	C1  string
}

type B3 struct {
	B1 A
	B2 int
}


func main() {
	a := A3{}
	fmt.Printf("a:%v\n", a)

}


