package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {
	host := "255.255.255.255"
	port := 3737
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%v", host, port))
	if err != nil {
		log.Panic(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	//_, err = conn.Read(data)
	//if err != nil {
	//	log.Panic(err)
	//}

	//t := binary.BigEndian.Uint32(data)
	//log.Println(time.Unix(int64(t), 0).String())
	go write(conn)
	time.Sleep(time.Millisecond * 10)
	for {
		data := make([]byte, 40)
		log.Println("---------")
		_, addr1, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Panic(err)
		}
		log.Println(string(data), addr1)
	}
	//os.Exit(0)

}

func write(conn *net.UDPConn) {
	for {
		_, err := conn.Write([]byte("111"))
		if err != nil {
			log.Panic(err)
		}
		log.Println("=========================")
		time.Sleep(30 * time.Second)
	}
}
