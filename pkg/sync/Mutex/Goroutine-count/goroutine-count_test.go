package Goroutine_count

import (
	"fmt"
	"testing"
	"time"
)

func TestCount(t *testing.T) {
	var mu Mutex
	for i := 0; i < 1000; i++ { // 启动1000个goroutine
		go func() {
			mu.Lock()
			time.Sleep(time.Millisecond)
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second)
	// 输出锁的信息
	fmt.Printf("waitings: %d, isLocked: %t, woken: %t,  starving: %t\n", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}
