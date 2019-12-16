package main

import "log"

func main() {
	log.Println("")
	if true {
		log.Println(1)
	}
	if false {
		log.Println(2)
	}
	if 1 == 1 {
		log.Println(3)
	}
}
