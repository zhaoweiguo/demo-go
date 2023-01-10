package main

import (
	"github.com/getsentry/sentry-go"
	"log"
	"runtime/debug"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	defer func() {
		time.Sleep(time.Second * 10)
		if err := recover(); err != nil {
			sentry.CurrentHub().Recover(err)
			sentry.Flush(time.Second * 5)

			log.Println("recover")
			time.Sleep(time.Second * 7)
			// print panic trace
			//panic(err)
		}
	}()

	go do()
}

func do() {
	defer Recover()
	time.Sleep(time.Second)
	panic(111)
	//panic(errors.New("132"))
	time.Sleep(time.Minute)
}

func Recover() {
	p := recover()
	if p != nil {
		switch e := p.(type) {
		case error:
			log.Println("msg: ", e, "\nstack:\n", string(debug.Stack()))
		default:
			log.Println("msg: ", e, "\nstack:\n", string(debug.Stack()))
		}
	}
}
