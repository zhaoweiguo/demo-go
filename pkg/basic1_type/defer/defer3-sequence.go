package main

import (
	"log"
	"time"
)

/*
defer 函数参数的值是代码执行到这段代码时的值，而不是代码运行的值
*/
func main() {
	defer func(t time.Time) {
		log.Println(t.Format(time.RFC3339)) // 1
		log.Println(time.Since(t))
	}(time.Now())
	log.Println(time.Now().Format(time.RFC3339)) // 2
	time.Sleep(time.Second * 2)
}
