package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	buf := []byte("Hello, 世界")
	fmt.Println("bytes =", len(buf))            // 13
	fmt.Println("runes =", utf8.RuneCount(buf)) // 9

	str := "Hello, 世界"
	fmt.Println("bytes =", len(str))                    // 13
	fmt.Println("runes =", utf8.RuneCountInString(str)) // 9

	fmt.Println(utf8.RuneLen('a')) // 1
	fmt.Println(utf8.RuneLen('界')) // 3

}
