package _map

import (
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
)

type PersonInfo struct {
	Id   string
	Name string
}

// 声明变量、赋值变量
func TestDeclare(t *testing.T) {
	// 先声明再赋值
	var a map[string]PersonInfo     // 变量声明
	a = make(map[string]PersonInfo) // 创建一个新map, ex: make(map[string] PersonInfo, 100)
	a["123"] = PersonInfo{"123", "gordon"}

	// 直接声明变量并赋值
	b := map[string]PersonInfo{
		"123": PersonInfo{"123", "gordon"},
	}
	assert.Equal(t, a, b)

	// 插入几条数据
	a["1"] = PersonInfo{"1", "simon"}
	fmt.Println(a)

}
