package main

import "C"
import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"unsafe"

	"github.com/3d0c/gmf"
	"github.com/hajimehoshi/oto"
)

func main() {
	// 创建音频播放器
	sampleRate := 48000
	channels := 2
	player, err := oto.NewPlayer(sampleRate, channels, 2, 2048)
	if err != nil {
		log.Fatal(err)
	}
	defer player.Close()

	// 创建 FFmpeg 解码器
	inCtx := gmf.NewCtx()
	defer gmf.Release(inCtx)
	if err := inCtx.OpenInput("rtsp://your-rtsp-url"); err != nil {
		log.Fatal(err)
	}
	if err := inCtx.FindStreamInfo(); err != nil {
		log.Fatal(err)
	}
	videoStream, err := inCtx.GetBestStream(gmf.AVMEDIA_TYPE_VIDEO)
	if err != nil {
		log.Fatal(err)
	}
	codec, err := gmf.FindDecoder(videoStream.CodecParameters().CodecID())
	if err != nil {
		log.Fatal(err)
	}
	defer gmf.Release(codec)
	if err := codec.Open(videoStream.CodecParameters()); err != nil {
		log.Fatal(err)
	}

	// 开始解码
	fmt.Println("Start decoding...")
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	frame := gmf.NewFrame()
	defer gmf.Release(frame)
	packet := gmf.NewPacket()
	defer gmf.Release(packet)
	for {
		select {
		case <-signalChan:
			fmt.Println("Interrupted!")
			return
		default:
			// 读取包并解码
			if err := inCtx.ReadPacket(packet); err != nil {
				log.Fatal(err)
			}
			if packet.StreamIndex() != videoStream.Index() {
				packet.Free()
				continue
			}
			if err := codec.Decode(packet.Data(), frame); err != nil {
				log.Fatal(err)
			}
			packet.Free()
			// 播放音频
			data, _, _ := frame.Samples().Data()
			dataBytes := C.GoBytes(unsafe.Pointer(data), C.int(frame.Samples().Size()))
			if _, err := player.Write(dataBytes); err != nil {
				log.Fatal(err)
			}
		}
	}
}
