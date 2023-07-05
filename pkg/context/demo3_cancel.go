package main

import (
	"context"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	go sub(cancel)
	select {
	case <-ctx.Done():
		log.Println("1")
	}
}

func sub(cancel context.CancelFunc) {
	time.Sleep(3 * time.Second)
	cancel()
}
