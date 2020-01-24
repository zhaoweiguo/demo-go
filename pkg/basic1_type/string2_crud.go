package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	demo1_string2()
}

// 字符串拼接
func demo1_string2() {
	str1 := "Gordon is a nice guy, isn't he?"
	str2 := " Yes, he is!"

	// 方法1: 运算符
	strA := str1 + str2
	fmt.Println(strA)

	// 方法2: fmt.Sprintf
	strB := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(strB)

	// 方法3: strings.Join
	strC := strings.Join([]string{str1, str2}, "->")
	fmt.Println(strC)

	// 方法4: 使用bytes.Buffer
	var buf bytes.Buffer
	buf.WriteString(str1)
	buf.WriteString(str2)
	fmt.Println(buf.String())

	// 方法5: 使用strings.Builder
	var build strings.Builder
	build.WriteString(str1)
	build.WriteString(str2)
	fmt.Println(build.String()) // good boy!

}
