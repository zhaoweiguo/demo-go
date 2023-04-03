package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"

	"github.com/use-go/onvif"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {
	searchCamera()
}

func newCamera(addr string) {
	log.Println(addr)
	param := onvif.DeviceParams{
		Xaddr:    "192.168.124.58:8899",
		Username: "admin",
		Password: "admin123",
	}
	dev, err := onvif.NewDevice(param)
	log.Println(dev, err)
	newCall(dev)
}

func newCall(device *onvif.Device) {
	data := GetDeviceInformation{}
	resp, err := device.CallMethod(data)
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

func searchCamera() error {
	re := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\d{1,5}\b`)
	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}
	for _, i := range interfaces {
		//log.Println(i)
		if i.Flags&net.FlagLoopback == net.FlagLoopback {
			continue
		}
		devices, err := onvif.GetAvailableDevicesAtSpecificEthernetInterface(i.Name)
		if err != nil {
			//log.Println(err)
			continue
		}
		log.Println("=================")
		for _, d := range devices {
			log.Println(d)
			log.Println(d.GetEndpoint("device"))
			addr := re.FindString(d.GetEndpoint("device"))
			newCamera(addr)
		}
	}
	return nil
}
