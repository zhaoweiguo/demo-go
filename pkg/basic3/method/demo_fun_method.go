package main

import (
	"fmt"
)

// 类型为func的method
// 注: 这块有点绕，值的多看几遍
// 注2: 此代码是我参考http模块的: type HandlerFunc func(ResponseWriter, *Request) 编造的

type DemoFunc func(int, string)

func (f DemoFunc) Help(a int, b string) {
	f(a, b)
}

type Helper interface {
	Help(int, string)
}

type HelpA struct{}

func (helpA HelpA) Help(a int, b string) {
	fmt.Println(a, b)
}

var f, h Helper

func main() {
	h = HelpA{}
	aaa := new(AAA)
	f = aaa.do(h)
	for i := 0; i < 10; i++ {
		f.Help(i, "gordon")
	}
}

type AAA struct {
	times int
}

func (aaa AAA) do(helper Helper) Helper {
	return DemoFunc(func(a int, b string) {
		fmt.Println(aaa.times)
		aaa.times++
		helper.Help(a, b)
	})
}
