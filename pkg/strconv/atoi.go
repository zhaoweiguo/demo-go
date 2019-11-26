package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
$ go build atoi.go
$ ./atoi 1 2 3 4
*/
func main() {
	args := os.Args
	v, error := strconv.Atoi(args[0]) // 把第一个参数转化为int

	fmt.Println(v, error)
}
