package _map

import (
	"fmt"
	"github.com/bmizerany/assert"
	"log"
	"testing"
)

// 从map查找
func TestIfExistString(t *testing.T) {
	maps := map[string]string{"key1": "value1"}

	// 可以赋值为两个变量
	newvalue, correct := maps["key1"]
	assert.Equal(t, newvalue, "value1")
	assert.Equal(t, correct, true)

	// 两个变量的，第二变量用于判断是否匹配成功存在
	wrongvalue, wrong := maps["wrongkey"]
	assert.Equal(t, wrongvalue, "")
	assert.Equal(t, wrong, false)

	wrongkey2 := maps["wrongkey"]
	assert.Equal(t, wrongkey2, "")
}

func TestIfExistBool(t *testing.T) {
	maps := map[string]bool{"key1": true}

	// 可以赋值为两个变量
	newvalue, correct := maps["key1"]
	assert.Equal(t, newvalue, true)
	assert.Equal(t, correct, true)

	// 两个变量的，第二变量用于判断是否匹配成功存在
	wrongvalue, wrong := maps["wrongkey"]
	assert.Equal(t, wrongvalue, false)
	assert.Equal(t, wrong, false)

	wrongkey2 := maps["wrongkey"]
	assert.Equal(t, wrongkey2, false)
}

type justForTest struct{}

func TestIfExistStruct(t *testing.T) {
	maps := map[string]justForTest{"key1": justForTest{}}

	// 可以赋值为两个变量
	newvalue, correct := maps["key1"]
	assert.Equal(t, newvalue, justForTest{})
	assert.Equal(t, correct, true)

	// 两个变量的，第二变量用于判断是否匹配成功存在
	wrongvalue, wrong := maps["wrongkey"]
	assert.Equal(t, wrongvalue, justForTest{})
	assert.Equal(t, wrong, false)

	wrongkey2 := maps["wrongkey"]
	assert.Equal(t, wrongkey2, justForTest{})
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
