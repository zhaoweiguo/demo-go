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
	fmt.Println(len(asciiString)) //输出：5
	fmt.Println(len(utf8String))  //输出：9
	// 获取字符长度
	fmt.Println(utf8.RuneCountInString(asciiString)) //输出：5
	fmt.Println(utf8.RuneCountInString(utf8String))  //输出：3
}

func TestDemo2(t *testing.T) {
	encodeRune()
	decodeRune()
}

// rune -> []byte
func encodeRune() {
	var r rune = 0x4E2D
	fmt.Printf("the unicode charactor is %c\n", r) // 中
	buf := make([]byte, 3)
	_ = utf8.EncodeRune(buf, r)                       // 对 rune 进行 utf-8 编码
	fmt.Printf("utf-8 representation is 0x%X\n", buf) // 0xE4B8AD
}

// []byte -> rune
func decodeRune() {
	var buf = []byte{0xE4, 0xB8, 0xAD}
	r, _ := utf8.DecodeRune(buf)
	fmt.Printf("the unicode charactor after decoding [0xE4, 0xB8, 0xAD] is %c\n", r)         // 中                                                        // 对 buf 进行 utf-8 解码
	fmt.Printf("the unicode charactor after decoding [0xE4, 0xB8, 0xAD] is %s\n", string(r)) // 中
}
