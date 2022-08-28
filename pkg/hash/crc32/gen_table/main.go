package main

import (
	"hash/crc32"
	"log"
)

func init()  {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {
	crc32_q()
	crc32_IEEE()
}

func crc32_q()  {
	// x³²+ x³¹+ x²⁴+ x²²+ x¹⁶+ x¹⁴+ x⁸+ x⁷+ x⁵+ x³+ x¹+ x⁰
	// reverse notation: 0b11010101100000101000001010000001
	// => 0xD5828281
	crc32q := crc32.MakeTable(0xD5828281)
	log.Printf("%08x\\n", crc32q)
	log.Printf("%d\\n", crc32q)
}

func crc32_IEEE()  {
	// X32+X26+X23+X22+X16+X12+X11+X10+X8+X7+X5+X4+X2+X1+X0
	// reverse notation: 0b11101101101110001000001100100000
	// 0x04C11DB7
	crc32q := crc32.MakeTable(crc32.IEEE)
	log.Printf("%08x\\n", crc32q)
	log.Printf("%d\\n", crc32q)
}