package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	cmdArguments := []string{"1000000000"}

	cmd := exec.Command("sleep", cmdArguments...)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cmd.Process)
	log.Println(cmd.ProcessState)
	log.Printf("command output: %q", out.String())
	go kill(cmd.Process)
	//go quit(cmd.Process)
	go cmd.Wait()
	select {}
}

func kill(proc *os.Process) {
	time.Sleep(5 * time.Second)
	proc.Kill()
}

func quit(proc *os.Process) {
	time.Sleep(5 * time.Second)
	proc.Signal(os.Interrupt)
}
