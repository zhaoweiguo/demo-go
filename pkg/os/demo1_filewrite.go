package main

import(
	"fmt"
	"os"
)

func main () {
	fmt.Println("start")
	// os.O_WRONLY: 只写, os.O_RDONLY只写, os.O_RDWR:读写
	// os.O_APPEND: 可与上面共用
	// os.O_CREATE: 没有时创建
	f, err := os.OpenFile("/tmp/abc.ttt", os.O_WRONLY | os.O_CREATE, 0600)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	if r, err := f.WriteString("1234\n"); err != nil {
		panic(err.Error())
		fmt.Println(r)
	}
	if r, err := f.WriteString("abcd\n"); err != nil {
		panic(err.Error())
		fmt.Println(r)
	}
}
