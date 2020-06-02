package json

import (
	"encoding/json"
	"log"
	"testing"
)

type JsonInterface struct {
	Result  interface{} `json:"result"`
	Value   int64       `json:"value"`
	Success bool        `json:"success"`
	T       int64       `json:"t"`
}

type JsonInterfaces struct {
	Result  []interface{} `json:"result"`
	Value   int64         `json:"value"`
	Success bool          `json:"success"`
	T       int64         `json:"t"`
}

type JsonMap struct {
	Result  []map[string]int64 `json:"result"`
	Value   int64              `json:"value"`
	Success bool               `json:"success"`
	T       int64              `json:"t"`
}

func TestUnmarshal1(t *testing.T) {

	var jsonBlob2 = []byte(`{"result":[{"20200503":2},{"20200504":8},{"20200505":4}], "value":123,"success":true,"t":1591081319805}`)
	var r1 JsonInterface
	json.Unmarshal(jsonBlob2, &r1)
	log.Println(r1)

	var r2 JsonInterfaces
	json.Unmarshal(jsonBlob2, &r2)
	log.Println(r2)

	var r3 JsonMap
	json.Unmarshal(jsonBlob2, &r3)
	log.Println(r3)
}

func TestUnmarshal2(t *testing.T) {

	var jsonBlob2 = []byte(`{"result":[{"20200503":2},{"20200504":8},{"20200505":4},{"20200506":4},{"20200507":8},{"20200508":4},{"20200509":3},{"20200510":1},{"20200511":13},{"20200512":4},{"20200513":2},{"20200514":8},{"20200515":28},{"20200516":1},{"20200517":1},{"20200518":37},{"20200519":31},{"20200520":26},{"20200521":8},{"20200522":2},{"20200523":1},{"20200524":2},{"20200525":2},{"20200526":2},{"20200527":1},{"20200528":0},{"20200529":2},{"20200530":3},{"20200531":2},{"20200601":0}],"success":true,"t":1591081319805}`)
	var r1 JsonInterface
	json.Unmarshal(jsonBlob2, &r1)
	log.Println(r1)

	var r2 JsonMap
	json.Unmarshal(jsonBlob2, &r2)
	log.Println(r2)
}
