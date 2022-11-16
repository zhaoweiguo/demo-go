package main

import (
	"log"
	"net"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {
	//local address
	la, err := net.ResolveUDPAddr("udp4", "0.0.0.0:37380")
	if err != nil {
		panic(err)
	}
	// remote address
	ra, err := net.ResolveUDPAddr("udp", "120.26.164.111:3737")
	if err != nil {
		panic(err)
	}

	// listen
	c, err := net.ListenUDP("udp4", la)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	for i := 0; i < 10; i++ {
		// write(指定 addr 的 write)
		n, err := c.WriteTo([]byte("111"), ra)
		if err != nil {
			log.Panic(err)
		}
		log.Println("writing start=========================", n)
		log.Println("RemoteAddr: ", c.RemoteAddr())
		log.Println("LocalAddr: ", c.LocalAddr())
		//log.Println("writing end=========================", n)

		// read
		data := make([]byte, 4096)
		n, addr, err := c.ReadFromUDP(data)
		log.Println("reading start=========================", n)
		log.Println(n, err, "[", addr, "]> message: ", string(data))
		//log.Println("reading end=========================", n)

		time.Sleep(time.Second * 5)
	}

	// As a Server: read
	log.Println("Server reading is start ... =========================")
	data := make([]byte, 4096)
	n, addr, err := c.ReadFromUDP(data)
	log.Println(n, addr, err)
	log.Println(string(data))
}
