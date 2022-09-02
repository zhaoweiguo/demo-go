package main

/**
1. init()只执行一次，只有在第一次使用此模块时被调用
*/

import (
	"github.com/zhaoweiguo/demo-go/pkg/basic3/init2/sub"
	"log"
)

func init() {
	log.Println("main init")
}

func main() {
	sub.Hello()
	sub.Hello()
}
