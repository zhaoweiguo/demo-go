package main

import (
	"log"

	"github.com/use-go/onvif"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {
	dev := auth1()
	log.Println(dev)
	log.Println(dev.GetServices())
	log.Println(dev.GetDeviceInfo())
}

func auth1() *onvif.Device {
	param := onvif.DeviceParams{
		Xaddr:    "192.168.124.58:8899",
		Username: "admin",
		Password: "admin123",
	}
	dev, err := onvif.NewDevice(param)
	if err != nil {
		panic(err)
	}
	return dev
}
