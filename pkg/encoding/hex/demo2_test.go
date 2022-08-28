package hex

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestD(t *testing.T)  {
	a := "AB"
	b, err := hex.DecodeString(a)
	if err != nil {
		panic(err)
	}
	log.Printf("%X", string(b))
	log.Println(string(b))
}



