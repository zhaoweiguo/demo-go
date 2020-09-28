package main

import "log"

/*
同一个包的init执行顺序，golang没有明确定义，编程时要注意程序不要依赖这个执行顺序
*/

const mainConst = 1

var mainVar = 2

func init() {
	log.Println("main.init")
	log.Println(subConst)
}

func main() {
	log.Println("main.main")
	subFunc()
}
