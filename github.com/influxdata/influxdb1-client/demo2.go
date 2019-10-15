package main

import (
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	client "github.com/influxdata/influxdb1-client"
)

func main()  {
	uri := "https://ts-2ze2r7ub895f9n48b.influxdata.rds.aliyuncs.com:3242"
	user := "test"
	passwd := "Q1w2e3r4"
	host, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}
	con, err := client.NewClient(client.Config{
		URL: *host,
		Username: user,
		Password: passwd,
	})
	if err != nil {
		log.Fatal(err)
	}

	var (
		shapes     = []string{"circle", "rectangle", "square", "triangle"}
		colors     = []string{"red", "blue", "green"}
		sampleSize = 1000
		pts        = make([]client.Point, sampleSize)
	)

	rand.Seed(42)
	for i := 0; i < sampleSize; i++ {
		pts[i] = client.Point{
			Measurement: "shapes",
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
		Database:        "test",
		RetentionPolicy: "autogen",
	}

	_, err = con.Write(bps)
	if err != nil {
		log.Fatal(err)
	}




}


