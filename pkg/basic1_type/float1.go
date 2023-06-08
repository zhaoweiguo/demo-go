package main

import (
	"log"
)

func main() {
	float1()
}

func float1() {
	// 这个为什么是相等的，是因为转换成二进制的数据是相等的
	var f1 float32 = 16777216.0
	var f2 float32 = 16777217.0
	log.Println(f1 == f2) // true
	log.Println(f1, f2)
}
