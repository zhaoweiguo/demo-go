package main

import (
	"fmt"
	"time"

	"github.com/giorgisio/goav/format"
	"github.com/giorgisio/goav/protocol/rtsp"
)

func main() {
	// RTSP URL
	url := "rtsp://username:password@host:port/stream_name"

	// 输出文件
	outUrl := "output.mp4"

	// 打开RTSP流
	input, err := rtsp.Open(url)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// 获取流信息
	info := input.Streams()
	dur := int64(info[0].Duration())
	fmt.Println(dur)

	// 创建输出上下文和流
	output, err := format.NewOutput(outUrl, input)
	if err != nil {
		panic(err)
	}
	stream := output.NewStream()

	// 复制流信息和参数
	stream.SetCodecParameters(info[0].CodecPar())
	stream.Duration(dur)

	// 启动转码
	if err := output.Start(); err != nil {
		panic(err)
	}
	<-time.After(5 * time.Second)
	output.Close()
}
