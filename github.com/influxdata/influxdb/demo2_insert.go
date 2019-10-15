package main

import (
	"fmt"
	"github.com/influxdata/influxdb/client"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

func main()  {
	host1 := "ts-bp16723l95rwey7a2.influxdata.rds.aliyuncs.com"
	port1 := 3242
	host, err := url.Parse(fmt.Sprintf("https://%s:%d", host1, port1))
	if err != nil {
		log.Fatal(err)
	}
	con, err := client.NewClient(client.Config{
		URL: *host,
		Username: "gordon",
		Password: "1QAZ2wsx",
	})
	if err != nil {
		log.Fatal(err)
	}
	var (
		shapes     = []string{"circle", "rectangle", "square", "triangle"}
		colors     = []string{"red", "blue", "green"}
		sampleSize = 100
		pts        = make([]client.Point, sampleSize)
	)

	rand.Seed(42)
	for i := 0; i < sampleSize; i++ {
		pts[i] = client.Point{
			Measurement: "user_line",
			Tags: map[string]string{
				"color": strconv.Itoa(rand.Intn(len(colors))),
				"shape": strconv.Itoa(rand.Intn(len(shapes))),
			},
			Fields: map[string]interface{}{
				"value": rand.Intn(sampleSize),
			},
			Time:      time.Now(),
			Precision: "s",
		}
	}

	bps := client.BatchPoints{
		Points:          pts,
		Database:        "dev",
		//RetentionPolicy: "default",
	}

	_, err = con.Write(bps)
	if err != nil {
		log.Fatal(err)
	}




}

