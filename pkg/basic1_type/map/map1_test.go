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
	// struct
	mapa := make(map[string]StructA)
	log.Println(mapa)

	itema := mapa["notexist"]
	log.Println(itema)
	log.Println(itema.Column1)
	log.Println(itema.Column2)
	log.Println(itema.Column3)

	log.Println("====================================")

	// *struct
	mapb := make(map[string]*StructA)
	log.Println(mapb)

	itemb := mapa["notexist"]
	log.Println(itemb)
	log.Println(itemb.Column1)
	log.Println(itemb.Column2)
	log.Println(itemb.Column3)

}
