// 主要用于测试statsd的准确度
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {
	connectionString := "10.3.255.215:8125"
	conn, err := net.Dial("udp", connectionString)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	i := 0
	for true{
		i = i+1
		if i % 100 == 0 {
			time := time.Now().Unix()
			fmt.Printf("i:%d, time:%v\n", i, time)
		}
		send(conn)
		time.Sleep(time.Millisecond * time.Duration(10))   // 0.01s
		
	}
}

func send(conn net.Conn){
	sampledData := make(map[string]string)

	var sampleRate float32
	sampleRate = 1
	data := map[string]string{}
	delta := 1   // 每次请求几个
	stats1 := "test.num.test1"
	data[stats1] = fmt.Sprintf("%d|c", delta)


//	fmt.Printf("data:%v\n", data)
	if sampleRate < 1 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rNum := r.Float32()
		if rNum <= sampleRate {
			for stat, value := range data {
				sampledUpdateString := fmt.Sprintf("%s|@%f", value, sampleRate)
				fmt.Println(sampledUpdateString)
				sampledData[stat] = sampledUpdateString
			}
		}
	} else {
		sampledData = data
	}

//	fmt.Printf("sampleData:%v\n", sampledData)

	for k, v := range sampledData {
		update_string := fmt.Sprintf("%s:%s", k, v)
//		fmt.Printf("update_string:%v\n", update_string)
		_, err := fmt.Fprintf(conn, update_string)
		if err != nil {
			log.Println(err)
		}
	}

}

