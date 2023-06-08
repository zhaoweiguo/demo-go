package main

import (
	"log"
)

type T struct {
	K1 string
	K2 int
}

func (t T) Name(k string) int {
	log.Println(t)
	return 0
}

func main() {
	var t T
	f1 := t.Name
	log.Printf("%T\n", f1)
	f2 := T.Name
	log.Printf("%T\n", f2)
	f3 := (*T).Name
	log.Printf("%T\n", f3)
}
