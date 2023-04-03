package main

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
)

// @todo
func main() {
	path := "./pkg/crypto/x509/demo2_user/id_rsa"
	path2 := "./pkg/crypto/x509/demo2_user/pem"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	p, _ := pem.Decode(data)
	log.Println(string(p.Bytes))
	ioutil.WriteFile(path2, p.Bytes, os.ModePerm)
	pem, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	log.Println(pem.Size(), pem.Validate(), pem.Public())
}
