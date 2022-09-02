package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDemo(t *testing.T) {
	str := "abcdefghiHello,世界"
	fmt.Println("Utf-8遍历")
	for i := 0; i < len(str); i++ {
		ch := str[i]
		ctype := reflect.TypeOf(ch)
		fmt.Printf("%s ", ctype)
		fmt.Println(ch)
	}
	fmt.Println("Unicode遍历")
	for _, ch1 := range str {
		ctype := reflect.TypeOf(ch1)
		fmt.Printf("%s ", ctype)

		fmt.Println(ch1)
	}
}
