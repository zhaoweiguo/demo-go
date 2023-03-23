package main

import (
	"fmt"
	"os"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gorilla/websocket"
	"github.com/pion/rtp"
	"github.com/pion/rtp/codecs"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/ivfreader"
	"github.com/pion/webrtc/v3/pkg/media/oggreader"
	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"
	"github.com/pion/webrtc/v3/pkg/media/oggwriter"
)

/**
使用 Golang 和 FFmpeg 库进行 H.264 视频解码的示例代码
*/

const (
	defaultWebsocketEndpoint = "ws://localhost:8080"
)

func main() {
	// Read the input file and determine its MIME type
	inputFilename := os.Args[1]
	inputMimetype, err := mimetype.DetectFile(inputFilename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Input file MIME type: %s\n", inputMimetype.String())

	// Initialize the WebRTC engine
	webrtc.RegisterDefaultCodecs()

	// Create a new MediaEngine instance
	mediaEngine := webrtc.MediaEngine{}
	if err = mediaEngine.PopulateFromSDP(`v=0
o=- 1337 1337 IN IP4 0.0.0.0
s=-
t=0 0
a=ice-ufrag:foobarbaz
a=ice-pwd:foobarbazfoobarbazfoobarbaz
a=ice-options:trickle
a=fingerprint:sha-256 16:02:6C:81:F0:92:07:E4:81:1F:1A:9B:DC:F9:01:BF:8C:1E:2D:87:94:AB:71:F8:CA:9C:46:9B:43:81:C8:EC
a=setup:active
a=mid:data
a=sctpmap:5000 webrtc-datachannel 1024
`); err != nil {
		panic(err)
	}

	// Create a new PeerConnection instance
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		panic(err)
	}

	// Create a new Track instance
	var track *webrtc.TrackLocalStaticSample
	switch inputMimetype.String() {
	case "video/h264":
		track, err = webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{
			MIMEType:     "video/h264",
			PayloadType:  96,
			ClockRate:    90000,
			RTCPFeedback: nil,
		}, "video", "pion")
		if err != nil {
			panic(err)
		}

		// Start a new goroutine to read the input file and send H264 packets to the Track
		go func() {
			f, err := os.Open(inputFilename)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			// Initialize the H264 decoder
			decoder, err := codecs.NewH264Decoder()
			if err != nil {
				panic(err)
			}

			// Initialize the RTP packetizer
			packetizer := rtp.NewPacketizer(1400, 96, 0, 0, nil, decoder.RTPCodecCapability().MimeType)

			// Create a new
		}
	}
}
