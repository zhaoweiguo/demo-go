package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	demo2()
}

func demo1() {
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

func demo2() {

	type Attributes struct {
		AttrTime    string   `json:"attr_time,omitempty"`
		AttributeId string   `json:"attribute_id,omitempty"`
		Value       []string `json:"value,omitempty"`
	}

	type Request struct {
		Attributes   []Attributes
		GadgetId     string `json:"gadget_id,omitempty"`
		GadgetTypeId string `json:"gadget_type_id,omitempty"`
	}

	type OriginMsg struct {
		From       string  `json:"from,omitempty"`
		Mac        string  `json:"mac,omitempty"`
		MsgType    string  `json:"msg_type,omitempty"`
		Request    Request `json:"request,omitempty"`
		SequenceId string  `json:"sequence_id,omitempty"`
		Version    string  `json:"version,omitempty"`
	}

	type Context struct {
		UserId     string    `json:"user_id,omitempty"`
		SequenceId string    `json:"sequence_id,omitempty"`
		OriginMsg  OriginMsg `json:"origin_msg,omitempty"`
	}

	type UpdateGadgetAttr struct {
		GadgetId   string  `json:"gadget_id,omitempty"`
		CreateTime string  `json:"create_time,omitempty"`
		Context    Context `json:"context,omitempty"`
	}

	var updateGadgetAttr = UpdateGadgetAttr{}

	a := []byte(`{"gadget_id":"c990036c1bfa535b7bc0e","create_time":"1575355298","context":{"user_id":"593fafedb34561","sequence_id":"e188e0a9b7a88e9e58e601e01b85af7c","origin_msg":{"msg_type":"update_gadget_attr","from":"hub/3de7d0667e148b569b8cf","mac":"default_resource","version":"v2.0","sequence_id":"e188e0a9b7a88e9e58e601e01b85af7c","request":{"gadget_id":"c990036c1bfab81","gadget_type_id":"200054","attributes":[{"attribute_id":"0x00300011","attr_time":"1575355298","value":["1"]}]}}}}`)
	b := json.Unmarshal(a, &updateGadgetAttr)
	log.Println(b)
}
