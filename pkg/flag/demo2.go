package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// go run demo2.go param1 param2
func main() {

	cmdline := os.Args[1:]
	fmt.Printf("cmdline: %s\n", cmdline)

	flag.CommandLine.Parse(cmdline)
	n := flag.NArg() // 参数个数
	log.Printf("n:%d\n", n)
	args := flag.Args() // 参数列表
	log.Printf("args: %s\n", args)

}
