package main

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
)


func queryDB(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: db,
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


func query()  {
	con, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     host1,
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	q := client.Query{
		Command:  "select count(*) from user_line",
		Database: db,
	}
	if response, err := con.Query(q); err == nil && response.Error() == nil {
		log.Println(response.Results)
	}

}


