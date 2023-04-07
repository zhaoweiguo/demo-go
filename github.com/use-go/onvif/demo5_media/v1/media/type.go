package media

import (
	"github.com/use-go/onvif/xsd"
)

type GetStreamUri struct {
	XMLName      string         `xml:"GetStreamUri"`
	StreamSetup  StreamSetup    `xml:"StreamSetup"`
	ProfileToken ReferenceToken `xml:"ProfileToken"`
}

type StreamSetup struct {
	Stream    StreamType `xml:"Stream"`
	Transport Transport  `xml:"Transport"`
}
type StreamType xsd.String
type Transport struct {
	Protocol TransportProtocol `xml:"Protocol"`
	Tunnel   *Transport        `xml:"Tunnel"`
}
type TransportProtocol xsd.String
type ReferenceToken xsd.String
