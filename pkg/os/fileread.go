package main

import (
	"fmt"
	"os"
)

// folder相关
func main() {
	demo1()
}

func demo1() {
	fmt.Println("demo1 start=====================")

	// 读取文件与文件夹方法相同
	f, err := os.OpenFile("/tmp", os.O_RDONLY, 0600)

	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	fmt.Println(f.Name())

	// 取出本文件的stat
	fi, err := f.Stat()
	fmt.Println(fi.Name(), fi.IsDir(), fi.Size(), fi.Mode(), fi.ModTime())

	fmt.Println("1=====================")
	// -1为取出全部子文件名
	// 取出此文件夹下所有文件的stat
	dirs, err := f.Readdir(-1)
	if err != nil {
		panic(err.Error())
	}
	for _, dir := range dirs {
		fmt.Println(dir.Name(), dir.Mode(), dir.Size(), dir.IsDir())
	}

	fmt.Println("2=====================")
	// 取出文件夹下所有文件的文件名
	// 注: 使用了一次上面的Readdir函数，再使用Readdirnames就没有数据了
	names, err := f.Readdirnames(-1)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(names)
}

func demo2() {
	fmt.Println("demo2 start=====================")
	folder := "/tmp"
	fi, err := os.Stat(folder)
	fmt.Println(err, fi.IsDir(), fi.Name())

}
