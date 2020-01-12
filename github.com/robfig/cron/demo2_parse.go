package main

import (
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	a := "* 1 * 1 * *"
	schedule, err := cron.Parse(a)
	log.Println(schedule.Next(time.Now()), err) // 算出离当前时间最近的下一次执行时间
}
