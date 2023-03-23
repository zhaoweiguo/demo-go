package main

import (
	"github.com/pion/webrtc/v3"
	"image"
	"image/draw"
	"image/jpeg"
	"os"

	"github.com/gorilla/websocket"
	"github.com/pion/rtp"
	"github.com/pion/rtp/codecs"
)

func main() {
	// create WebRTC configuration
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	// create WebRTC API object
	api := webrtc.NewAPI(webrtc.WithMediaEngine(&webrtc.MediaEngine{}))

	// create a new PeerConnection object
	pc, err := api.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	// create a video track
	videoTrack, err := pc.NewTrack(webrtc.MimeTypeH264, 1, "video", "pion")
	if err != nil {
		panic(err)
	}

	// add the video track to the PeerConnection
	_, err = pc.AddTrack(videoTrack)
	if err != nil {
		panic(err)
	}

	// create a new WebSocket connection
	wsConn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		panic(err)
	}

	// receive RTP packets over WebSocket
	for {
		_, rtpPacket, err := wsConn.ReadMessage()
		if err != nil {
			panic(err)
		}

		// create a new RTP packet
		rtpPacketObj := &rtp.Packet{}
		err = rtpPacketObj.Unmarshal(rtpPacket)
		if err != nil {
			panic(err)
		}

		// create a new H264 decoder
		h264Decoder := codecs.NewH264Decoder()

		// decode the RTP packet into a video frame
		videoFrame, err := h264Decoder.Decode(rtpPacketObj.Payload)
		if err != nil {
			panic(err)
		}

		// create an image from the video frame
		ycbcrImage := image.NewYCbCr(image.Rect(0, 0, int(videoFrame.Width), int(videoFrame.Height)), image.YCbCrSubsampleRatio420)
		copy(ycbcrImage.Y, videoFrame.Y)
		copy(ycbcrImage.Cb, videoFrame.Cb)
		copy(ycbcrImage.Cr, videoFrame.Cr)
		rgbaImage := image.NewRGBA(ycbcrImage.Bounds())
		draw.Draw(rgbaImage, rgbaImage.Bounds(), &image.YCbCr{Y: ycbcrImage.Y, Cb: ycbcrImage.Cb, Cr: ycbcrImage.Cr, YStride: ycbcrImage.YStride, CStride: ycbcrImage.CStride, SubsampleRatio: ycbcrImage.SubsampleRatio}, image.Point{0, 0}, draw.Src)

		// write the image to a JPEG file
		jpegFile, err := os.Create("frame.jpg")
		if err != nil {
			panic(err)
		}
		defer jpegFile.Close()
		jpeg.Encode(jpegFile, rgbaImage, &jpeg.Options{Quality: 90})
	}
}
