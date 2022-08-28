package basic_encode

import (
	"log"
	"testing"
)

func TestAbc(t *testing.T)  {
	// 所有这些编码方式中，0—127 表示的符号是一样的
	log.Println("\u004A")	// J
	log.Println("\u007E")	// ~
	log.Printf("%X", "\u004A")	// 4A
	log.Printf("%X", "\u007E")	// 7E

	// 不同的编码方式，128—255 的这一段不一样
	log.Println("\u0082")	//
	log.Println("\u00AB")	// «
	log.Println("\u00CD")	// Í
	log.Println("\u00EF")	// ï
	log.Printf("%X", "\u0082")	// C282
	log.Printf("%X", "\u00AB")	// C2AB
	log.Printf("%X", "\u00CD")	// C38D
	log.Printf("%X", "\u00EF")	// C3AF

}
