package main

import (
	"log"
	"runtime"
)

func main() {
	demo1_Caller()
	demo2_Callers()
	demo3()
}

func demo1_Caller() {
	log.Println("==demo1 start===================================")
	pc, file, line, ok := runtime.Caller(0)
	log.Println("value:", pc, file, line, ok)

	pc, file, line, ok = runtime.Caller(1)
	log.Println("value:", pc, file, line, ok)

	pc, file, line, ok = runtime.Caller(2)
	log.Println("value:", pc, file, line, ok)

	pc, file, line, ok = runtime.Caller(3)
	log.Println("value:", pc, file, line, ok)

	pc, file, line, ok = runtime.Caller(4)
	log.Println("value:", pc, file, line, ok)
	log.Println("==demo1 end===================================")
}

func demo2_Callers() {
	log.Println("==demo2_user start===================================")
	pcs := make([]uintptr, 2)
	_ = runtime.Callers(0, pcs)
	fun := runtime.FuncForPC(pcs[1])
	log.Println(fun.Name(), fun.Entry())
	log.Println("==demo2_user end===================================")
}

func demo3() {
	log.Println("==demo3 start===================================")
	pcs := make([]uintptr, 25)
	depth := runtime.Callers(4, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	log.Println(depth, "=", frames)
	for f, again := frames.Next(); again; f, again = frames.Next() {
		log.Println("-")
		log.Println(f.Function, f.Func, f.Entry, f.File, f.Line, f.PC)
	}
	log.Println("==demo3 end===================================")
}
