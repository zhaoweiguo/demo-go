package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	normal1()
	normal2_env()
	normal3()
}

func normal1() {
	fmt.Println("normal1 start=====================")
	homedir, err := os.UserHomeDir()
	fmt.Println(homedir, err)

	configdir, err := os.UserConfigDir()
	fmt.Println(configdir)

	cachedir, err := os.UserCacheDir()
	fmt.Println(cachedir)

	tmpdir := os.TempDir()
	fmt.Println(tmpdir)

	link := ""
	os.Readlink(link)

	os.Lstat("/tmp")

}

func normal2_env() {
	fmt.Println("normal2_env start=====================")
	env := os.Getenv("HOME")
	fmt.Println(env)
}

func normal3() {
	fmt.Println("normal3 文件(夹)操作 start=====================")

	testRootfolder := "./1111111"
	os.Mkdir(testRootfolder, os.ModePerm)
	abc1, err := os.Create(testRootfolder + "/abc1")
	log.Println(abc1, err)
	os.Create(testRootfolder + "/abc2")
	os.Create(testRootfolder + "/abc3")

	// 注: 需要注释掉才能显示
	os.Remove(testRootfolder + "/abc1")

	// 注: 需要注释掉才能看到上面创建的
	os.RemoveAll(testRootfolder)

}
