package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	a = "a/c/b"
	b := strings.Split(a, "/")      // b
	c := strings.HasPrefix(a, "a/") // true
	fmt.Println(b, c)
}
