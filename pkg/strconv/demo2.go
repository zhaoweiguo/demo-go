package main

import (
"encoding/json"
"fmt"
"strconv"
)

type Msg struct {
	Channel string
	Name string
	Msg string
}

func main() {
	var msg Msg
	var val = []byte(`"{\"channel\":\"buu\",\"name\":\"john\", \"msg\":\"doe\"}"`)

	s, _ := strconv.Unquote(string(val))

	err := json.Unmarshal([]byte(s), &msg)

	fmt.Println(string(val))
	fmt.Println(s)
	fmt.Println(err)
	fmt.Println(msg.Channel, msg.Name, msg.Msg)



	var val2 = `{"channel":"buu","name":"john", "msg":"doe"}`
	err2 := json.Unmarshal([]byte(s), &msg)

	fmt.Println(val2)
	fmt.Println(err2)

}




