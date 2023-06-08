package xml

import (
	"encoding/xml"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestMediaUrl(t *testing.T) {
	do()
}

type GetStreamUri struct {
	XMLName      string `xml:"GetStreamUri"`
	StreamSetup  string `xml:"space StreamSetup"`
	ProfileToken string `xml:"ProfileToken"`
}

func do() {
	uri := GetStreamUri{
		XMLName:      "XMLName",
		StreamSetup:  "StreamSetup",
		ProfileToken: "Token",
	}
	output, err := xml.MarshalIndent(uri, "  ", "    ")
	log.Println(string(output), err)
}
