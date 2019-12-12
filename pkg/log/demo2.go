package main

import (
	"log"
)

func main() {
	log.Println(1)
	log.SetFlags(log.Ldate)
	log.Println(2)
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println(3)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println(4)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println(5)
	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime | log.Lmicroseconds)
	log.Println(6)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.LUTC | log.Lmicroseconds)
	log.Println(7)
}
