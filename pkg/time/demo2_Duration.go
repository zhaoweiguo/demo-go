package main

import(
	"fmt"
	"time"
)

func main() {
	timeout := time.Duration(1) * time.Second
	fmt.Printf("timeout: %d\n", timeout)   // 1000,000,000

	timeout = time.Duration(1) * time.Millisecond
	fmt.Printf("timeout: %d\n", timeout)   // 1000,000

	timeout = time.Duration(1) * time.Microsecond
	fmt.Printf("timeout: %d\n", timeout)   // 1000

	timeout = time.Duration(1) * time.Minute
	fmt.Printf("timeout: %d\n", timeout)   // 60,000,000,000




}
