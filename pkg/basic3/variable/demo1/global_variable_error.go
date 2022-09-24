package main

import "log"

/*
:= 赋值的变量都是局部变量，即使其中有全局变量也会忽略
*/

var x string // 全局变量1
var y string // 全局变量2
func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	//x, err := getValue()   // 这儿的x被:=改变为非全局变量
	var err error
	x, err = getValue() // 这儿的x还是全局变量
	log.Println(x, err)
}

func getValue() (string, error) {
	return "name", nil
}

func doit() {
	y, err := getValue()
	log.Println(y, err)
}
func doit2() {
	var err error
	y, err = getValue()
	log.Println(y, err)
}

func main() {
	log.Println(x)
	doit()
	log.Println(y)
	doit2()
	log.Println(y)
}
