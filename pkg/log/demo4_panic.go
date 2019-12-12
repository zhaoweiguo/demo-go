package main

import (
	"log"
)

func main() {
	// Panic is equivalent to Print() followed by a call to panic().
	log.Panicln("panic")
	log.Println("can not be logged")
}
