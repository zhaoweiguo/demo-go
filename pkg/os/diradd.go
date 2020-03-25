package main

import (
	"log"
	"os"
)

func main() {
	base := "./pkg/os/data/"
	file1 := "abc.txt"
	file2 := "a/b/c.txt"
	// 创建文件夹
	err := os.Mkdir(base+file1, os.ModePerm)
	log.Println(err)
	// 级联创建文件夹
	err = os.MkdirAll(base+file2, os.ModePerm)
	log.Println(err)
}
