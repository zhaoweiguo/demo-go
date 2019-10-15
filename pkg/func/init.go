package main

import "fmt"

func main() {
	fmt.Println("main")
}

func init() {
	fmt.Println("init")
}

func initAbc() {
	fmt.Println("initAbc")
}
