package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("runtime.Version:%s\n", runtime.Version())
	fmt.Printf("runtime.GOOS:%s\n", runtime.GOOS)
	fmt.Printf("runtime.GOARCH:%s\n", runtime.GOARCH)
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())
	fmt.Println("runtime.NumCPU()", runtime.NumCPU())
	fmt.Println("runtime.GOROOT()", runtime.GOROOT())
}
