package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	a = "a/c/b"
	b := strings.Split(a, "/")
	fmt.Println(b[len(b)-1])
}
