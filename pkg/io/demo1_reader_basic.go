package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/*
对于要用作读取器的类型，它必须实现 io.Reader 接口的唯一一个方法 Read(p []byte)。
换句话说，只要实现了 Read(p []byte) ，那它就是一个读取器。

type Reader interface {
    Read(p []byte) (n int, err error)
}

*/
func main() {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF:", n)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(n, string(p[:n]))
	}

}
