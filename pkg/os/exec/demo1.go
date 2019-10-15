package main

import (
	"os/exec"
	"strings"
	"bytes"
	"fmt"
	"log"
)

func main() {
	// shell命令: tr a-z A-Z
	cmd := exec.Command("tr", "a-z", "A-Z")
	// 输入
	cmd.Stdin = strings.NewReader("some input")
	// 输出
	var out bytes.Buffer
	cmd.Stdout = &out
	// 运行此shell
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())

}