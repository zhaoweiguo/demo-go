package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	count int32
	wg    sync.WaitGroup
)

/*
基于basic3/concurrency/demo1.go项目改进
增加:
atomic.LoadInt32(&count), atomic.StoreInt32(&count,value)
一个读取int32类型变量的值，一个是修改int32类型变量的值，这两个都是原子性的操作
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
		value := atomic.LoadInt32(&count)
		runtime.Gosched()
		value++
		atomic.StoreInt32(&count, value)
	}
}
