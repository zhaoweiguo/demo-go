package main

import (
	_ "github.com/zhaoweiguo/demo-go/pkg/basic3/import/multiLevelDir/sub2"

	"github.com/zhaoweiguo/demo-go/pkg/basic3/import/multiLevelDir/sub"
	_ "github.com/zhaoweiguo/demo-go/pkg/basic3/import/multiLevelDir/sub1"
	"log"
)

/*
初始化顺序：变量初始化->init()->main()

不同包的init函数按照包导入的依赖关系决定执行顺序，即：
多级package执行时，
1. 先执行最里层init函数
2. 再执行次里层init函数
...
3. 直到执行main的init函数
4. 最后再执行main函数
*/
func init() {
	log.Println("main.init")
}

func main() {
	log.Println("main.main")
	sub.SubFunc()
}
