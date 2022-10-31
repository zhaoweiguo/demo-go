package main

import (
	"flag"
	"log"
	"net"
	"time"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "3738", "port")

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		log.Panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		log.Panic(err)
	}
	log.Println(n, remoteAddr)
	log.Println(string(data))

	daytime := time.Now().String()
	log.Println(daytime)
	//b := make([]byte, 40)
	//binary.BigEndian.PutUint32(b, uint32(daytime))
	conn.WriteToUDP([]byte(daytime), remoteAddr)
}
