package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/use-go/onvif"
	"github.com/zhaoweiguo/demo-go/github.com/use-go/onvif/demo5_media/v1/media"
	"github.com/zhaoweiguo/demo-go/github.com/use-go/onvif/utils"
)

/*
说明: 这儿直接自定义onvif的结构体，有问题
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

func newCall(dev *onvif.Device) {

	data := media.GetServiceCapabilities{} //
	//data := getStreamUri()

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

func getStreamUri() media.GetStreamUri {
	return media.GetStreamUri{
		StreamSetup: media.StreamSetup{
			Stream: utils.Stream,
			Transport: media.Transport{
				Protocol: utils.Protocol,
				Tunnel:   nil,
			},
		},
		ProfileToken: utils.Token,
	}
}
