package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, _ := os.Open("pkg/bufio/aaa.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		log.Println(scanner.Text())
	}

}
