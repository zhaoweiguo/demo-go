package main

import "time"

func main() {
	doit1()
}

func doit1() {
	date := time.Now()
	if date.Format("MST") == "UTC" {
		date = date.Add(time.Hour * 8)
	}

}
