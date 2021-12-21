package int

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func TestInt(t *testing.T) {
	a := 100 * 23 / 79
	log.Printf("a: %d\n", a)
}
