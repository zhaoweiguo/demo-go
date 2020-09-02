package main

import (
	"log"
	"runtime"
	"sync"
)

/*
sync.Once.Do(f func())能保证once只执行一次,这个sync.Once块只会执行一次。
*/
var callerInitOnce sync.Once

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, 2)
		_ = runtime.Callers(0, pcs)
		logrusPackage := runtime.FuncForPC(pcs[1]).Name()
		log.Println(logrusPackage)
	})
}
