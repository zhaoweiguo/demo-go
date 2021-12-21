package main

import (
	"github.com/zhaoweiguo/demo-go/pkg/basic3/concurrency/demo3_daemon/work"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	go work.Run()
	for i := 0; i < 10; i++ {
		log.Println(work.Get())
		time.Sleep(time.Second)
	}
}
