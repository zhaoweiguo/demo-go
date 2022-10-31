package main

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
)

var (
	url = "http://localhost:8086"
	org = "myorg"
	bucket = "analyzer"
	token = "SJE321PSuVxBiBSH0bp83pOIeUSs1emYO_8xv-lbjDhIRGKv0XigvAppftIyG2KXbI00ZBQAe4k951wO0X1j0g=="
)


func main() {

	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient(url, token)
	write(client)
	//query(client)
	// Ensures background processes finishes
	client.Close()
}

func write(client influxdb2.Client)  {
	// Use blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking(org, bucket)
	write1(writeAPI)
	//write2(writeAPI)
	//write3(writeAPI)

}
func write1(writeAPI api.WriteAPIBlocking)  {
	// Create point using full params constructor
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45.0},
		time.Now())
	// write point immediately
	if err := writeAPI.WritePoint(context.Background(), p); err != nil {
		panic(err.Error())
	}
}
func write2(writeAPI api.WriteAPIBlocking)  {
	// Create point using fluent style
	p := influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 23.2).
		AddField("max", 45.0).
		SetTime(time.Now())
	if err := writeAPI.WritePoint(context.Background(), p); err != nil {
		panic(err.Error())
	}
}
func write3(writeAPI api.WriteAPIBlocking)  {
	// Or write directly line protocol
	line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0)
	err := writeAPI.WriteRecord(context.Background(), line)
	if err != nil {
		panic(err.Error())
	}
}


func query(client influxdb2.Client)  {
	// Get query client
	queryAPI := client.QueryAPI(org)
	// Get parser flux query result
	result, err := queryAPI.Query(context.Background(), `from(bucket:"analyzer")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
	if err != nil {
		panic(err.Error())
	}

	// Use Next() to iterate over query result lines
	for result.Next() {
		// Observe when there is new grouping key producing new table
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		// read result
		fmt.Printf("row: %s\n", result.Record().String())
	}
	if result.Err() != nil {
		fmt.Printf("Query error: %s\n", result.Err().Error())
	}
}
