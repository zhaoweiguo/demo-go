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
		in.WriteString("ffmpeg -rtsp_transport tcp -i rtsp://admin:admin123@192.168.124.2:554/live/ch0 -t 5 -c:v libx264 -c:a aac /home/xinxi/test/output3.mp4\n")
		in.WriteString("exit\n")
	}()
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
