package utf8

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestDemo1(t *testing.T) {
	asciiString := "abcde"
	utf8String := "中国人"
	//输出字节长度
	fmt.Println(len(asciiString))	//输出：5
	fmt.Println(len(utf8String))	//输出：9
	// 获取字符长度
	fmt.Println(utf8.RuneCountInString(asciiString)) //输出：5
	fmt.Println(utf8.RuneCountInString(utf8String))	 //输出：3
}



