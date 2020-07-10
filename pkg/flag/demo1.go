package main

import (
	"flag"
	"fmt"
)

// go run demo1_simple.go --param1 param1 --param2 param2
func main() {
	var (
		param1 = flag.String("param1", "1111", "参数1")
		param2 = flag.String("param2", "value2", "参数2")
	)
	// 解析前为默认值
	fmt.Println(">> 解析前")
	fmt.Printf("param1:%s\n", *param1)
	fmt.Printf("param2:%s\n", *param2)

	// 解析
	flag.Parse()

	fmt.Println(">> 解析后")
	fmt.Printf("param1:%s\n", *param1)
	fmt.Printf("param2:%s\n", *param2)

	fmt.Println("打印帮助.........")
	flag.PrintDefaults()

}
