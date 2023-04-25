package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	in := bytes.NewBuffer(nil)
	//linux
	//cmd := exec.Command("sh")
	//windows
	cmd := exec.Command("sh")
	cmd.Stdin = in
	go func() {
		in.WriteString("echo hello world >> ./os/exec/output/test.txt\n")
		in.WriteString("exit\n")
	}()
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
