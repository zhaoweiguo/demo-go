package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/use-go/onvif"
	"github.com/use-go/onvif/media"
	onvid "github.com/use-go/onvif/xsd/onvif"
	"github.com/zhaoweiguo/demo-go/github.com/use-go/onvif/utils"
)

/*
说明: 这儿直接使用onvif的结构体，有问题
*/

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {
	param := onvif.DeviceParams{
		Xaddr:    utils.Addr,
		Username: utils.Username,
		Password: utils.Password,
	}
	dev, err := onvif.NewDevice(param)
	log.Println(dev, err)
	newCall(dev)
}

func getStreamUri() media.GetStreamUri {
	return media.GetStreamUri{
		StreamSetup: onvid.StreamSetup{
			Stream: utils.Stream,
			Transport: onvid.Transport{
				Protocol: utils.Protocol,
				//Protocol: "UDP",
				Tunnel: nil,
			},
		},
		ProfileToken: utils.Token,
	}
}

func newCall(dev *onvif.Device) {
	//data := media.GetMetadataConfigurationOptions{}
	//data := media.GetVideoSources{}
	//data := media.GetStreamUri{}

	//data := media.GetVideoEncoderConfiguration{ConfigurationToken: onvid.ReferenceToken("ve000")}
	//data := media.GetVideoEncoderConfigurations{}

	//data := media.GetServiceCapabilities{} // Method 'trt:GetServiceCapabilities' not implemented: method name or namespace not recognized

	//data := media.GetProfiles{}
	data := getStreamUri() // ter:InvalidArgVal
	//data := media.GetSnapshotUri{ProfileToken: onvid.ReferenceToken("profile_000")} // ter:InvalidArgVal

	//data := media.GetAudioSources{}

	resp, err := dev.CallMethod(data)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("======:", resp.StatusCode)
		body, err := ioutil.ReadAll(resp.Body)
		log.Println(string(body), err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(body))
	return

}
