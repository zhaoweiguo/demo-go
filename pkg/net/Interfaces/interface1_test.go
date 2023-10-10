package Interfaces

import (
	"log"
	"net"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestInterfaceAddr(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
		return
	}

	//for _, addr := range addrs {
	//	log.Println("===============>")
	//	log.Println(addr)
	//	log.Println("===============<")
	//}
	log.Println("=============================")
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.Println(ipnet.String())
				log.Println(ipnet.Network())
				return
			}
		}
	}
}
func TestInterface(t *testing.T) {

	addrs, err := net.Interfaces()

	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		if addr.Flags&net.FlagLoopback != 0 {
			continue
		}
		if addr.Flags&net.FlagRunning == 0 {
			continue
		}
		subaddrs, _ := addr.Addrs()
		//log.Println("===============>")
		//log.Println(addr.Addrs())
		//log.Println("===============<")
		for _, subaddr := range subaddrs {
			if ipnet, ok := subaddr.(*net.IPNet); ok {
				if ipnet.IP.To4() != nil {
					log.Println("-------")
					log.Println(ipnet)
					log.Println(addr.Addrs())
					log.Println(addr.Name, addr.HardwareAddr, addr.Index, addr.MTU)
					log.Println(addr.Flags)
					log.Println()
					log.Println(addr.MulticastAddrs())
					return // 第一个获取到的非Loopback, 运行
				}
			}
		}
	}
}
