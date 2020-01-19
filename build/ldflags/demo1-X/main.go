package main

import (
	"fmt"
	"os"
	"runtime"
)

// 在编译后的二进制文件中注入 git 等版本信息后再发布
/*
编译的时候，通过链接选项 -X 来动态传入版本信息
flags="-X 'main.goversion=$(go version)' -X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash=`git describe --tags --always --dirty`"
go build -ldflags "$flags" -x -o demo1 main.go
./demo1 -v
*/
var buildstamp = ""
var githash = ""
var goversion = fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)

func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		// 注: 这儿打印的是编译时的git commit hash， 编译时间， 编译机上的golang版本
		fmt.Printf("Git Commit Hash: %s\n", githash)
		fmt.Printf("UTC Build Time : %s\n", buildstamp)
		fmt.Printf("Golang Version : %s\n", goversion)

		return
	}
}
