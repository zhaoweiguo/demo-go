package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)


func main() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGALRM, syscall.SIGBUS, syscall.SIGCHLD, syscall.SIGCONT, syscall.SIGEMT, syscall.SIGFPE, syscall.SIGILL, syscall.SIGINFO, syscall.SIGHUP, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGKILL)   // 监听这两种signal

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)


}


