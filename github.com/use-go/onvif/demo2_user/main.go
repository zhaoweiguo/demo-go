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
		Xaddr:      "192.168.124.58:8899",
		Username:   "admin",
		Password:   "admin123",
		HttpClient: new(http.Client),
	}
	dev, err := onvif.NewDevice(param)
	log.Println(dev, err)

	data := device.GetUsers{}
	res, err := dev.CallMethod(data)
	bs, _ := ioutil.ReadAll(res.Body)
	log.Printf("output %+v %s", res.StatusCode, bs)

}
