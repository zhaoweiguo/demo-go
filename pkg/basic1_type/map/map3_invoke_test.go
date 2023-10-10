package _map

import (
	"fmt"
	"testing"
)

func TestDemo_map1(t *testing.T) {
	// 函数间传递Map是不会拷贝一个该Map的副本的，
	//也就是说如果一个Map传递给一个函数，该函数对这个Map做了修改，那么这个Map的所有引用，都会感知到这个修改
	dict := map[string]int{"王五": 60, "张三": 43}
	modify_map(dict)
	fmt.Println(dict["张三"])
}

func modify_map(dict map[string]int) {
	dict["张三"] = 10
}
