package main

/*
from: https://www.flysnow.org/2017/04/29/go-in-action-go-runner.html
*/
import (
	"log"
	"os"
	"time"

	"github.com/zhaoweiguo/demo-go/pkg/os/signal/demo2/common"
)

func main() {
	log.Println("...开始执行任务...")

	timeout := 3 * time.Second
	r := common.New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case common.ErrTimeOut:
			log.Println(err)
			os.Exit(1)
		case common.ErrInterrupt:
			log.Println(err)
			os.Exit(2)
		}
	}
	log.Println("...任务执行结束...")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("正在执行任务%d", id)
		time.Sleep(time.Duration(id) * time.Second)
		log.Printf("完成任务%d", id)
	}
}
