package main

import (
	"log"
	"os"
)

func main() {
	str, _ := os.Getwd()
	log.Println(str)
}
