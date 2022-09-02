package hex

import (
	"fmt"
	"log"
	"testing"
)

func TestInt(t *testing.T) {
	var a uint32
	a = 4294967295
	strA := fmt.Sprintf("%X", a) // FFFFFFFF
	log.Println(strA)

	var b uint32
	b = 1
	strB := fmt.Sprintf("%08X", b) // FFFFFFFF
	log.Println(strB)
}
