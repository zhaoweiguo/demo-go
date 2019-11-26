package main

import (
	"fmt"
	"os"
)

// folder相关
func main() {
	fmt.Println("start")

	// 读取文件与文件夹方法相同
	f, err := os.OpenFile("/tmp", os.O_RDONLY, 0600)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	// -1为取出全部子文件名
	names, err := f.Readdirnames(-1)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(names)
}
