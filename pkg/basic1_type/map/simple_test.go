package _map

import (
	"fmt"
	"github.com/bmizerany/assert"
	"log"
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

// 从map查找
func TestNotExist(t *testing.T) {
	maps := map[string]string{"key1": "value1"}

	// 可以赋值为两个变量
	newkey, correct := maps["key1"]
	assert.Equal(t, newkey, "value1")
	assert.Equal(t, correct, true)

	// 两个变量的，第二变量用于判断是否匹配成功存在
	wrongkey, wrong := maps["wrongkey"]
	assert.Equal(t, wrongkey, "")
	assert.Equal(t, wrong, false)
}

// 从map查找
func TestFind(t *testing.T) {
	var maps = make(map[string]string)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		maps[key] = value
	}

	log.Printf("maps:%v\n", maps) // [key1=>value1 key2=>value2...]

	// 也可以赋值为一个变量
	onekey := maps["key1"]
	fmt.Printf("onekey: %v\n", onekey)
	// 如一个变量，匹配不成功的情况
	notfound := maps["wrongkey"]
	fmt.Printf("notfound: %v\n", notfound)
}
