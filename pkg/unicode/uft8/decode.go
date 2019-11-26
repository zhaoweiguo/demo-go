package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// example1
	b := []byte("Hello, 世界")
	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		fmt.Printf("%c %v|", r, size) // 界 3|世 3|  1|, 1|o 1|l 1|l 1|e 1|H 1|
		b = b[:len(b)-size]
	}
	fmt.Println("")

	// example2
	str := "Hello, 世界"
	for len(str) > 0 {
		r, size := utf8.DecodeLastRuneInString(str)
		fmt.Printf("%c %v|", r, size) // 界 3|世 3|  1|, 1|o 1|l 1|l 1|e 1|H 1|
		str = str[:len(str)-size]
	}
	fmt.Println("")

	// example3
	b2 := []byte("Hello, 世界")
	for len(b2) > 0 {
		r, size := utf8.DecodeRune(b2)
		fmt.Printf("%c %v|", r, size) // H 1|e 1|l 1|l 1|o 1|, 1|  1|世 3|界 3|

		b2 = b2[size:]
	}
	fmt.Println("")

	// example4
	str2 := "Hello, 世界"
	for len(str2) > 0 {
		r, size := utf8.DecodeRuneInString(str2)
		fmt.Printf("%c %v|", r, size) // H 1|e 1|l 1|l 1|o 1|, 1|  1|世 3|界 3|
		str2 = str2[size:]
	}
	fmt.Println("")

	// example5
	r := '世'
	buf := make([]byte, 3)
	n := utf8.EncodeRune(buf, r)
	fmt.Println(buf) // [228 184 150]
	fmt.Println(n)   // 3
	fmt.Println("")

}
