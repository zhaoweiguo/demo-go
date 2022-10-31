package main

import (
	"github.com/influxdata/influxdb/client"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)



func insert()  {
	host, err := url.Parse(host1)
	if err != nil {
		log.Fatal(err)
	}
	con, err := client.NewClient(client.Config{
		URL: *host,
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
		Database:        db,
		//RetentionPolicy: "default",
	}

	_, err = con.Write(bps)
	if err != nil {
		log.Fatal(err)
	}




}

