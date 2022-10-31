package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

var host = flag.String("host", "0.0.0.0", "host")
var port = flag.String("port", "3737", "port")

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

	//host := "255.255.255.255"
	//port := 3736
	//addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%v", host, port))

	packet := time.Now().Format(time.RFC3339)
	n, err = conn.WriteTo([]byte(packet), remoteAddr)
	broadCast("255.255.255.255", remoteAddr.Port, packet)
}

func broadCast(host string, port int, message string) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%v", host, port))
	if err != nil {
		log.Panic(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	for {
		_, err := conn.Write([]byte(message))
		if err != nil {
			log.Panic(err)
		}
		log.Println("=========================")
		time.Sleep(30 * time.Second)
	}
}
