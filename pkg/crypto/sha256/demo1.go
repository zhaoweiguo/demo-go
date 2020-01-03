package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

/*
sha256是golang内置的sha256的散列标准库，可以让我们很容易的生成对应数据的散列值
*/
func main() {
	demo1()
	demo2()
}

func demo1() {
	fmt.Println("===demo1============================")
	// 第一种调用方法
	sum := sha256.Sum256([]byte("hello world\n"))
	fmt.Printf("%x\n", sum)

	// 第二种调用方法
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	fmt.Printf("%x\n", h.Sum(nil))
}

// 对文件加密
func demo2() {
	fmt.Println("===demo2============================")
	f, err := os.Open("pkg/crypto/sha256/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x\n", h.Sum(nil))
}
