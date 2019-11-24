package main

import (
	"flag"
	"fmt"
)

// 使用命令说明:
// go run demo3.go -port=8000 -port2=7000 -i abc.html enable remote

func main() {
	var port int

	var infile *string = flag.String("i", "defaultfile", "File contains values for sorting")
	var port2 *int = flag.Int("port2", 88, "other port")
	fmt.Println(*port2)

	flag.IntVar(&port, "port", 80, "help for port")

	fmt.Printf("%s\n", *infile) // defaultfile
	fmt.Println(port)           // 80
	fmt.Println(flag.Args())    //[]

	// !Action: notice the different after Parse function
	flag.Parse()

	fmt.Println(*infile) // abc.html
	fmt.Println(port)    //8000
	fmt.Println(*port2)  //8000

	fmt.Println(flag.Args()) //[enable remote]
	fmt.Println(flag.Arg(0)) // enable
	fmt.Println(flag.Arg(1)) // remote
}
