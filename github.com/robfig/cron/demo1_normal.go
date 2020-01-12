package main

import (
	"github.com/robfig/cron"
	"log"
)

func main() {
	demo1()
}

func demo1() {
	c := cron.New() // 新建一个定时任务对象
	i := 1
	c.AddFunc("*/1 * * * * *", func() { // 1秒执行一次
		log.Println(i)
		i++
	}) // 给对象增加定时任务
	c.Start()
	select {}
}
