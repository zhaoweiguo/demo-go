package Interfaces

import (
	"log"
	"net"
	"testing"
)

func TestInterface(t *testing.T) {

	addrs, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		log.Println("-------")
		log.Println(addr.Name, addr.HardwareAddr, addr.Index, addr.Flags, addr.MTU)
		log.Println(addr.MulticastAddrs())
		log.Println(addr.Addrs())
	}
}
