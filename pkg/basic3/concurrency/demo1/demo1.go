package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int32
	wg    sync.WaitGroup
)

/*
这是一个资源竞争的例子:
我们可以多运行几次这个程序，会发现结果可能是2，也可以是3，也可能是4。
因为共享资源count变量没有任何同步保护，所以两个goroutine都会对其进行读写，会导致对已经计算好的结果覆盖，以至于产生错误结果

共享资源竞争的问题，非常复杂，并且难以察觉，好在Go为我们提供了一个工具帮助我们检查，这个就是:
go build -race命令
我们在当前项目目录下执行这个命令，生成一个可以执行文件，然后再运行这个可执行文件，就可以看到打印出的检测信息
*/
func main() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched()
		value++
		count = value
	}
}
