package main

import (
	"log"
)

func main() {
	var a, b, c string
	a = "a"
	b = "b"
	c = "c"
	src := []string{a, b, c}
	log.Println(src)
	for index, value := range src {
		log.Println(index, value)
	}

}
