package main

import (
	"context"
	"log"
	"time"
)

/*
超时处理实例
*/

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		log.Println("[main] context超时:", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		log.Println("[handle]context超时:", ctx.Err()) // 注意: 有时能执行到，有时执行不到。执行不到是因为main进程已经结束。
	case <-time.After(duration):
		log.Println("处理完成，用时", duration)
	}
}
