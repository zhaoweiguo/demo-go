package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "3737", "port")

var sessions map[string]*net.UDPAddr
var uids []string

type Data struct {
	Uid  string `json:"uid"`
	Data string `json:"data"`
}

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

	sessions = make(map[string]*net.UDPAddr, 10)
	uids = []string{"gordon", "simon", "bland"}

	go sendMsg(conn)
	for {
		handleClient(conn)
	}
}

func sendMsg(conn *net.UDPConn) {
	for {
		for _, uid := range uids {
			log.Println(uid)
			log.Println(sessions)
			addr, correct := sessions[uid]
			if !correct {
				log.Println("-->", uid)
			} else {
				log.Println("==>", uid)
				conn.WriteTo([]byte(fmt.Sprintf("recived: %s", uid)), addr)
			}
			time.Sleep(2 * time.Second)
		}
		log.Println("======")
		time.Sleep(5 * time.Second)
	}
}

func handleClient(conn *net.UDPConn) {
	data2 := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data2)
	data := data2[:n]
	if err != nil {
		log.Panic(err)
	}
	log.Println(n, remoteAddr)
	log.Println(fmt.Sprintf("%s:--%v", string(data), cap(data)))
	var model Data
	err = json.Unmarshal(data, &model)
	if err != nil {
		log.Panic(err)
	}
	log.Println(model)
	uid := model.Uid
	sessions[uid] = remoteAddr // 每次client上报都更新

	daytime := time.Now().String()
	log.Println(daytime)
	//b := make([]byte, 40)
	//binary.BigEndian.PutUint32(b, uint32(daytime))
	conn.WriteToUDP([]byte(daytime), remoteAddr)
}
