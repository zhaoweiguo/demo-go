package main

import (
	"fmt"
)
func main() {

	type Site map[string]string

	var maps = make(Site)
	for i:=0; i<10; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		maps[key] = value
	}

	fmt.Printf("maps:%v\n", maps)  // [key1=>value1 key2=>value2...]

	// 可以赋值为两个变量
	newkey, correct := maps["key1"]
	fmt.Printf("newkey: %s, correct: %t\n", newkey, correct)   // Note: correct:true

	// 也可以赋值为一个变量
	onekey := maps["key1"]
	fmt.Printf("onekey: %v\n", onekey)

	// 两个变量的，第二变量用于判断是否匹配成功存在
	wrongkey, wrong := maps["wrongkey"]
	fmt.Printf("wrongkey: %s, wrong: %t\n", wrongkey, wrong)  // Note: wrong:false

	// 如一个变量，匹配不成功的情况
	notfound := maps["wrongkey"]
	fmt.Printf("onekey not found things: %v\n", notfound)
}
