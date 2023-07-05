package main

import (
	"bytes"
	"context"
	"log"
	"os/exec"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	cmdArguments := []string{"1000000000"}

	parentCtx := context.Background()
	ctx, Cancel := context.WithCancel(parentCtx)
	cmd := exec.CommandContext(ctx, "sleep", cmdArguments...)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("command output: %q", out.String())
	go kill(Cancel)
	cmd.Wait()
}

func kill(cancel context.CancelFunc) {
	time.Sleep(5 * time.Second)
	cancel()
}
