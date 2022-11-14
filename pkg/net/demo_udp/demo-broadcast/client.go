package main

import (
	"fmt"
	"log"
	"net"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {
	//local address
	la, err := net.ResolveUDPAddr("udp4", "0.0.0.0:3738")
	ra, err := net.ResolveUDPAddr("udp", "192.168.95.255:3737")

	// listen
	// 需要使用 ListenUDP 而不是 DialUDP。 使用 DialUDP 时，它会创建一个 "已连接" 的 UDP 端口，并且仅在读取时返回源自远程地址的数据包。
	// https://www.codenong.com/40897417/
	c, err := net.ListenUDP("udp4", la)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	// write(指定 addr 的 write)
	n, err := c.WriteTo([]byte("111"), ra)
	if err != nil {
		log.Panic(err)
	}
	log.Println("=========================", n)

	// read
	data := make([]byte, 40)
	log.Println("RemoteAddr: ", c.RemoteAddr())
	log.Println("LocalAddr: ", c.LocalAddr())
	log.Println(c.ReadFrom(data))
	log.Println(string(data))
}
