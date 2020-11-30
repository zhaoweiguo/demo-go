package json

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDemoMap(t *testing.T) {
	var jsonBlob2 = []byte(`{"user_id":"1", "create_time":"123", "context":{"name":"dddd"}}`)
	var mapBody map[string]interface{}
	err := json.Unmarshal(jsonBlob2, &mapBody)
	log.Println(err)
	log.Println(mapBody)

}
