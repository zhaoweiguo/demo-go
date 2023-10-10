package _map

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Llongfile | log.Ltime)
}

type StructA struct {
	Column1 string
	Column2 int
	Column3 map[string]int
}

func TestMake(t *testing.T) {
	mapa := make(map[string]StructA)
	log.Println(mapa)

	itema := mapa["notexist"]
	log.Println(itema)
}
