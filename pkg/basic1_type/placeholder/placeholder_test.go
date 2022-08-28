package placeholder

import (
	"log"
	"testing"
)

func init()  {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

// 整数占位符:
/*
%b        二进制表示                   Printf("%b", 5)               101
%c        Unicode 码点所表字符          Printf("%c", 0x4E2D)          中
%d        十进制表示                   Printf("%d", 0x12)            18
%o        八进制表示                   Printf("%o", 10)              12
%q        单引号围绕的字符字面值         Printf("%q", 0x4E2D)         '中'
%x        十六进制,字母小写: a-f        Printf("%x", 13)               d
%x        十六进制,字母小写: a-f        Printf("%08x", 13)             0000000d
%X        十六进制,字母大写: A-F        Printf("%X", 13)               D
%U        Unicode 格式:               Printf("%U", 0x4E2D)  U+4E2D
          U+1234，等同于 "U+%04X"
*/
func TestInt(t *testing.T)  {
	log.Printf("%b", 5)				// 101
	log.Printf("%c", 0x4E2D)			// 中
	log.Printf("%d", 0x12)			// 18
	log.Printf("%o", 10)				// 12
	log.Printf("%q", 0x4E2D)			// '中'
	log.Printf("%x", 13)				// d
	log.Printf("%08x", 13)           // 0000000d
	log.Printf("%X", 13)				// D
	log.Printf("%U", 0x4E2D)  		// U+4E2D
}












