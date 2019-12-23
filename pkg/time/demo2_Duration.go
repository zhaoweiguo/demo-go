package main

import (
	"fmt"
	"time"
)

func main() {

	timeout := time.Duration(1) * time.Microsecond
	time.Sleep(timeout)
	fmt.Printf("timeout: %d\n", timeout) // 1000

	timeout = time.Duration(1) * time.Millisecond
	time.Sleep(timeout)
	fmt.Printf("timeout: %d\n", timeout) // 1000,000

	timeout = time.Duration(1) * time.Second
	time.Sleep(timeout)
	fmt.Printf("timeout: %d\n", timeout) // 1000,000,000

	timeout = time.Duration(1) * time.Minute
	time.Sleep(timeout)
	fmt.Printf("timeout: %d\n", timeout) // 60,000,000,000

}
