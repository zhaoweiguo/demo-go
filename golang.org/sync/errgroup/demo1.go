package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

// 【注】只返回第一个错误
func main() {
	var g errgroup.Group
	// 启动第一个子任务,它执行成功
	g.Go(func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("exec #1")
		return nil
	})
	// 启动第二个子任务，它执行失败
	g.Go(func() error {
		time.Sleep(3 * time.Second)
		fmt.Println("exec #2")
		return errors.New("failed to exec #2")
	})
	// 启动第三个子任务，它执行成功
	g.Go(func() error {
		time.Sleep(8 * time.Second)
		fmt.Println("exec #3")
		return nil
	})
	// 启动第4个子任务，它执行失败
	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #4")
		return errors.New("failed to exec #4")
	})

	// 等待4个任务都完成
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully exec all")
	} else {
		fmt.Println("failed:", err)
	}
}
