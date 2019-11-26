package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	buf := []byte{228, 184, 150}        // 世
	fmt.Println(utf8.FullRune(buf))     // true
	fmt.Println(utf8.FullRune(buf[:2])) // false
	fmt.Println("----")

	str3 := "世"
	fmt.Println(utf8.FullRuneInString(str3))     // true
	fmt.Println(utf8.FullRuneInString(str3[:2])) // false
	fmt.Println("----")

	buf2 := []byte("a界")
	fmt.Println(utf8.RuneStart(buf2[0])) // true
	fmt.Println(utf8.RuneStart(buf2[1])) // true
	fmt.Println(utf8.RuneStart(buf2[2])) // false
	fmt.Println("----")

	valid := []byte("Hello, 世界")
	invalid := []byte{0xff, 0xfe, 0xfd}
	fmt.Println(utf8.Valid(valid))   // true
	fmt.Println(utf8.Valid(invalid)) // false
	fmt.Println("----")

	valid2 := 'a'
	invalid2 := rune(0xfffffff)
	fmt.Println(utf8.ValidRune(valid2))   // true
	fmt.Println(utf8.ValidRune(invalid2)) // false
	fmt.Println("----")

	valid3 := "Hello, 世界"
	invalid3 := string([]byte{0xff, 0xfe, 0xfd})
	fmt.Println(utf8.ValidString(valid3))   // true
	fmt.Println(utf8.ValidString(invalid3)) // false

}
