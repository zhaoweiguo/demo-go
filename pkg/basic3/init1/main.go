package main

import (
	"log"

	"github.com/zhaoweiguo/demo-go/pkg/basic3/init1/sub1"
)

func init() {
	log.Println("main init")
}

func main() {
	sub1.Hello12()
	sub1.Hello()
}
