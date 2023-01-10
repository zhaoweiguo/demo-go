package main

import (
	"log"
	"net"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {

	//// 指定client侧端口
	//la, err := net.ResolveUDPAddr("udp4", "0.0.0.0:3737")
	//if err != nil {
	//	panic(err)
	//}
	//c, err := net.ListenUDP("udp4", la)
	c, err := net.ListenUDP("udp4", nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// write(指定 addr 的 write)
	ra, err := net.ResolveUDPAddr("udp", "223.104.38.195:35151")
	n, err := c.WriteToUDP([]byte("222"), ra)
	if err != nil {
		log.Panic(err)
	}
	log.Println("=========================", n)
	log.Println("RemoteAddr: ", c.RemoteAddr())
	log.Println("LocalAddr: ", c.LocalAddr())
}
