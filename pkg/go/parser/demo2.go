package main

import (
	"go/token"
	"fmt"
	"go/parser"
	"path/filepath"
)

func main() {
	fmt.Println("start1>")
	token := token.NewFileSet()
	fmt.Printf("[1]: %T => %v\n", token, token)

	// 找到指定目录下以.go结尾的文件名列表
	pkgs, err := parser.ParseDir(token, filepath.Join(".", "pkg", "flag"), nil, parser.PackageClauseOnly)
	if err != nil {
		panic(err.Error())
	}
	for name := range pkgs{
		// 打印文件名(去除.go)
		fmt.Println(name)
	}


}

