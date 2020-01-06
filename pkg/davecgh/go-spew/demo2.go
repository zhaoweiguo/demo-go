package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	i := 1
	fmt.Println(spew.Sdump(i))

	s := "str"
	fmt.Println(spew.Sdump(s))

	ss := []string{"aa", "bb", "cc"}
	fmt.Println(spew.Sdump(ss))

	zwg := gordon{"zhaowg", true, 18}
	fmt.Println(spew.Sdump(zwg))
	spew.Dump(zwg) // 与上面一句等同
}

type gordon struct {
	name string
	sex  bool
	age  int
}
