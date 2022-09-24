package main

import (
	"github.com/zhaoweiguo/demo-go/pkg/basic3/variable/demo_global/global"
	"log"
)

func main() {
	log.Println(global.Get())
	global.Do1()
	log.Println(global.Get())
	global.Do2()
	log.Println(global.Get())
	global.Do3()
	log.Println(global.Get())
}
