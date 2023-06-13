package main

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	and()
}

func and() {
	a := 0x20
	log.Println(a) // 32
	log.Println(a & 0x1F)
}
