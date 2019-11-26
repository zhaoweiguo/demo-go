package main

import (
	"fmt"
	"os"
)

// $ go build args.go
// $ ./args a b c d
func main() {
	args := os.Args
	fmt.Println(args)
	fmt.Println(len(args))
}
