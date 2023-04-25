package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmdArguments := []string{"-rtsp_transport", "tcp",
		"-i", "rtsp://admin:admin123@192.168.124.2:554/live/ch0",
		"-c:v", "libx264",
		"-c:a", "aac", "/home/xinxi/test/output4.mp4"}

	cmd := exec.Command("ffmpeg", cmdArguments...)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("command output: %q", out.String())
}
