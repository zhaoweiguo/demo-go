package s

import (
	"fmt"
	"testing"
	"time"
)

func TestDemo1(t *testing.T) {
	var ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
			fmt.Println(i, ",")
		case v := <-ch:
			fmt.Println(v)
		}
	}
	time.Sleep(time.Second * 2)
}
