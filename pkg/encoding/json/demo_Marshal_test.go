package json

import (
	"encoding/json"
	"fmt"
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

// We'll use these two structs to demonstrate encoding and
// decoding of custom types below.
type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type response3 struct {
	page   int
	fruits []string
}

func TestSimpleBool(t *testing.T) {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
}

func TestSimpleInt(t *testing.T) {
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
}

func TestSimpleFloat(t *testing.T) {
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
}

func TestSimpleString(t *testing.T) {
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))
}

func TestSimpleArray(t *testing.T) {
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	// MarshalIndent就是数据格式化的Marshal功能
	slcBPretty, _ := json.MarshalIndent(slcD, "", "  ")
	fmt.Println(string(slcBPretty))
}

func TestSimpleMap(t *testing.T) {
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
}

func TestSimpleMap2(t *testing.T) {
	mapD := map[string]interface{}{"apple": 5, "lettuce": []string{"a", "b"}}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
}

func TestMarshalStruct1(t *testing.T) {
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res1BPretty, _ := json.MarshalIndent(res1D, "#", "   ")
	fmt.Println(string(res1BPretty))
}

// 指定json的字段名(与结构体response1不同的点)
func TestMarshalStruct2(t *testing.T) {
	// struct with tag
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	assert.Equal(t, []byte(`{"page":1,"fruits":["apple","peach","pear"]}`), res2B)
}

// 字段名为小写时不可见(与结构体response1不同的点)
func TestMarshalStruct3(t *testing.T) {
	// struct with tag
	res2D := &response3{
		page:   1,
		fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D) // 字段小写不可见,所以返回值为空
	assert.Equal(t, []byte(`{}`), res2B)
}

type responseComplex struct {
	Response1 response1
	Response2 response2
	values    string
}

func TestMarshalStruct4(t *testing.T) {
	// struct with tag
	res1 := response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2 := response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res3 := &responseComplex{
		Response1: res1,
		Response2: res2,
		values:    "vvvvvv",
	}
	res4, _ := json.Marshal(res3)
	fmt.Println(string(res4))
}

func TestDemo23_array(t *testing.T) {
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)
}

func TestStruct3(t *testing.T) {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
