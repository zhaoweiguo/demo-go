// go build -buildmode=plugin -o plugin.so
package main

import "fmt"

var V int

func F() { fmt.Printf("Hello, number %d\n", V) }
