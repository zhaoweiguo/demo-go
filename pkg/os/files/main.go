package main

import (
	"log"
	"os"
)

func main()  {
	log.Println(os.Getwd())

	f, err := os.Open("os/files/a.txt")
	if err != nil {
		panic(err)
	}
	log.Println(f.Name())
}


