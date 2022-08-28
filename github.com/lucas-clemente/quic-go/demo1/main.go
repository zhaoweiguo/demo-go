package main

import (
	"github.com/lucas-clemente/quic-go/http3"
	"net/http"
)

func main()  {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http3.ListenAndServeQUIC("localhost:4242", "./asserts/cert/server.crt", "./asserts/cert/server_key.pem", nil)
}
