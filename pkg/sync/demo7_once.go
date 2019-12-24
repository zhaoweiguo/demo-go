package main

import (
	"log"
	"runtime"
	"sync"
)

var callerInitOnce sync.Once

func main() {
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, 2)
		_ = runtime.Callers(0, pcs)
		logrusPackage := runtime.FuncForPC(pcs[1]).Name()
		log.Println(logrusPackage)
	})
}
