package main

import (
	"fmt"
	"github.com/gomedia/rtsp"
	"log"
)

func main() {
	// RTSP客户端连接配置
	rtspConfig := rtsp.ClientConfig{
		StreamURI: "rtsp://wowzaec2demo.streamlock.net/vod/mp4:BigBuckBunny_115k.mov",
		Codec:     "H264",
	}
	// 建立RTSP客户端连接
	client, err := rtsp.Dial(rtspConfig)
	if err != nil {
		log.Fatal(err)
	}
	// 获取RTSP连接的SDP信息
	sdp, err := client.Describe()
	if err != nil {
		log.Fatal(err)
	}
	// 获取RTSP流的Track信息
	tracks, err := sdp.GetTracks()
	if err != nil {
		log.Fatal(err)
	}
	// 打印Track信息
	for _, track := range tracks {
		fmt.Printf("Track %d: codec=%s\n", track.ID, track.Codec)
	}
	// 开始播放RTSP流
	err = client.Play()
	if err != nil {
		log.Fatal(err)
	}
	// 读取RTSP流数据
	for {
		pkt, err := client.ReadPacket()
		if err != nil {
			log.Fatal(err)
		}
		// 处理RTSP流数据
		fmt.Printf("Read packet: timestamp=%d, type=%s, size=%d\n", pkt.Timestamp, pkt.Type, len(pkt.Data))
	}
}
