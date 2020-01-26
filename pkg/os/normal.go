package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	normal1_common()
	normal2_env()
	normal3_op()
}

func normal1_common() {
	fmt.Println("normal1 start=====================")
	homedir, err := os.UserHomeDir()
	fmt.Println(homedir, err)

	configdir, err := os.UserConfigDir()
	fmt.Println(configdir)

	cachedir, err := os.UserCacheDir()
	fmt.Println(cachedir)

	tmpdir := os.TempDir()
	fmt.Println(tmpdir)

	linkPath := "./github/golang"
	link, err := os.Readlink(linkPath)
	fmt.Println(link)

	fi, err := os.Lstat("/tmp")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(fi.Name(), fi.IsDir())

}

func normal2_env() {
	fmt.Println("normal2_env start=====================")
	env := os.Getenv("HOME")
	fmt.Println(env)
}

// 文件(夹)相关操作
func normal3_op() {
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

	source := "./go.mod/"
	target := "/tmp/target"
	os.Link(source, target)    // 硬链接
	os.Symlink(source, target) // 软链接

}
