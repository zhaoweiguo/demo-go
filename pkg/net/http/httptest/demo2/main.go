package main

import (
	"encoding/json"
	"net/http"
)

func route() func(rw http.ResponseWriter, r *http.Request) {
	//API调用处理函数
	sendJson := func(rw http.ResponseWriter, r *http.Request) {
		u := struct {
			Name string
		}{
			Name: "张三",
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(u)
	}
	return sendJson
}
