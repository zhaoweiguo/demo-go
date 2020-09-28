package main

import "log"

const subConst = 1

var subVar = 2

func init() {
	log.Println("sub.init")
	log.Println(mainConst)
}

func subFunc() {
	log.Println("sub.f")
}
