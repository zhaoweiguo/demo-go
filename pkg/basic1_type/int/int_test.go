package int

import (
	"encoding/binary"
	"log"
	"testing"
	"unsafe"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func TestInt(t *testing.T) {
	a := 100 / 3
	log.Printf("a: %d\n", a) // 33

}

func TestLen(t *testing.T) {
	var a uint32 = 12
	log.Println(unsafe.Sizeof(a)) // 4
	log.Println(binary.Size(a))   // 4

}
