package main

import "fmt"

type PersonInfo struct {
	Id   string
	Name string
}

func main() {
	var personDB map[string]PersonInfo     // 变量声明
	personDB = make(map[string]PersonInfo) // 创建一个新map, ex: make(map[string] PersonInfo, 100)
	/*
	   1. 创建并初使化map代码如下:
	   map[string] PeronsInfo {
	       "1234" : PersonInfo{"1", "gordon"}
	   }
	*/

	// 插入几条数据
	personDB["1"] = PersonInfo{"1", "gordon"}
	personDB["123"] = PersonInfo{"123", "leo"}

	person, ok := personDB["111"]

	// ok是一个bool类型,true代表找到对应数据
	if ok {
		fmt.Println(person.Name)
	} else {
		fmt.Println("not found...")
	}

}
