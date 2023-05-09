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
		//in.WriteString("ffmpeg -rtsp_transport tcp -i rtsp://admin:admin123@192.168.124.2:554/live/ch0 -t 5 -c:v libx264 -c:a aac -f flv /home/xinxi/test/output1.flv\n")
		//in.WriteString("ffmpeg -rtsp_transport tcp -i rtsp://admin:admin123@192.168.124.2:554/live/ch0 -t 5 -c:v mpeg2video -c:a mp2 -f mpeg /home/xinxi/test/output2.mpg\n")
		//in.WriteString("ffmpeg -rtsp_transport tcp -i rtsp://admin:admin123@192.168.124.2:554/live/ch0 -t 5 -vcodec wmv2 /home/xinxi/test/output3.wmv\n")
		//in.WriteString("ffmpeg -rtsp_transport tcp -i rtsp://admin:admin123@192.168.124.2:554/live/ch0 -t 5 -c:v libx264 -c:a aac -f matroska /home/xinxi/test/output4.mkv\n")
		//in.WriteString("ffmpeg -rtsp_transport tcp -i rtsp://admin:admin123@192.168.124.2:554/live/ch0 -t 5 -c:v libx264 -c:a aac -f avi /home/xinxi/test/output5.avi\n")
		//in.WriteString("ffmpeg -rtsp_transport tcp -i rtsp://admin:admin123@192.168.124.2:554/live/ch0 -t 5 -c:v libx265 -c:a aac /home/xinxi/test/output6.mp4\n")
		//in.WriteString("ffmpeg -rtsp_transport tcp -i rtsp://admin:admin123@192.168.124.2:554/live/ch0 -t 5 -c:v libx264 -c:a aac -f mov /home/xinxi/test/output7.mov\n")
		in.WriteString("ffmpeg -rtsp_transport tcp -i rtsp://admin:admin123@192.168.124.2:554/live/ch0 -t 5 -c:v copy -c:a copy  /home/xinxi/test/output8.avi\n")
		in.WriteString("exit\n")
	}()
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
