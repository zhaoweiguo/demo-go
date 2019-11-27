package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var jsonBlob = []byte(`[
	{"na_me": "Pl2at_ypus", "Order": "Monostremata"},
	{"Na_me": "2wQuoll2",    "Order": "2Dasyuromorphia"}
]`)
	type Animal struct {
		Na_Me string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	log.Printf("%+v", animals)

	var jsonBlob2 = []byte(`{"user_id":"1", "create_time":"123", "context":{"name":"dddd"}}`)

	type Blob2 struct {
		User_Id     string
		Create_Time string
		Context     struct {
			Name string
		}
	}
	var blob2 Blob2
	err = json.Unmarshal(jsonBlob2, &blob2)
	if err != nil {
		log.Println(err)
	}
	log.Println(blob2)

}
