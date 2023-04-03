package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/use-go/onvif"
	"github.com/use-go/onvif/device"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {
	param := onvif.DeviceParams{
		Xaddr:    "192.168.124.58:8899",
		Username: "admin",
		Password: "admin123",
	}
	dev, err := onvif.NewDevice(param)
	log.Println(dev, err)
	newCall(dev)
}

func newCall(dev *onvif.Device) {
	//data := device.GetDeviceInformation{}
	//data := device.GetSystemDateAndTime{}
	data := device.GetCapabilities{Category: "All"}
	//data := ptz.GetServiceCapabilities{}
	resp, err := dev.CallMethod(data)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("======:", resp.StatusCode)
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
