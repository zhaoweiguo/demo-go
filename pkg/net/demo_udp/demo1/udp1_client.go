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
	host := "127.0.0.1"
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

	_, err = conn.Write([]byte(""))
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 40)
	//_, err = conn.Read(data)
	//if err != nil {
	//	log.Panic(err)
	//}

	//t := binary.BigEndian.Uint32(data)
	//log.Println(time.Unix(int64(t), 0).String())
	go read(conn)
	for {
		_, err = conn.Read(data)
		if err != nil {
			log.Panic(err)
		}
		log.Println(string(data))

	}
	//os.Exit(0)

}

func read(conn *net.UDPConn) {
	for {
		_, err := conn.Write([]byte("111"))
		if err != nil {
			log.Panic(err)
		}
		time.Sleep(10 * time.Second)
	}
}
