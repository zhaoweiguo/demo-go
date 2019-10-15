package main

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
)

const (
	MyDB = "dev"
	username = "gordon"
	password = "1QAZ2wsx"
)

func queryDB(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: MyDB,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}


func main()  {
	con, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "https://ts-bp16723l95rwey7a2.influxdata.rds.aliyuncs.com:3242",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	q := client.Query{
		Command:  "select count(*) from user_line2",
		Database: MyDB,
	}
	if response, err := con.Query(q); err == nil && response.Error() == nil {
		log.Println(response.Results)
	}



}


