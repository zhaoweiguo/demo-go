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
	a := 100 * 23 / 79
	log.Printf("a: %d\n", a)
}

func TestLen(t *testing.T) {
	var a uint32 = 12
	log.Println(unsafe.Sizeof(a))
	log.Println(binary.Size(a))

}
