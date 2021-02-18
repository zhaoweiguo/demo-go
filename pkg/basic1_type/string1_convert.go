package main

import (
	"fmt"
	"strconv"
)

func main() {
	convertAscII()
	stringToInt()
	intToString()
}

func stringToInt() {
	fmt.Println()
	fmt.Println("====start stringToInt()....")
	a := "12345"
	// @action: 都是以变量a字串为基准(int16类型和xx进制都是指变量a)，结果都是int64类型10进制
	b, _ := strconv.ParseInt(a, 10, 0)   // int类型10进制
	fmt.Printf("b:按10进制转[%T]%d\n", b, b) // 12345

	c, _ := strconv.ParseInt(a, 16, 8)  // 按int8类型16进制进行转
	fmt.Printf("c:按8进制转[%T]%d\n", c, c) //127 (本来结果是0X12345对应的10进制，但由于int8限制，剩下的截断了)

	fmt.Println("====end string to int....")
}

func intToString() {
	fmt.Println()
	fmt.Println("====start intToString()....")
	a := 12345
	b := fmt.Sprintf("str_%d", a)
	fmt.Println(b)
	fmt.Println("====end int to string....")
}

func convertAscII() {
	fmt.Println()
	fmt.Println("====start convertAscII()....")
	// 把ascII码转为字符
	a := 111
	b := string(a)
	if b == "o" {
		fmt.Printf("b:[%T]%s\n", b, b)
	}
	fmt.Println("====end convertAscII....")
}
