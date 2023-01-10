package main

import (
	"github.com/robfig/cron/v3"
	"log"
)

func main() {
	c := cron.New(cron.WithSeconds())
	i := 1
	c.AddFunc("*/1 * * * * *", func() { // 1秒执行一次
		log.Println(i)
		i++
	}) // 给对象增加定时任务
	c.Start()
	select {}
}
