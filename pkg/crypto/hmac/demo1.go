package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
)

func main() {
	secret := "SEC36b0e14e99a510e42444d8921cc142cf84926aa11ce2f42f9126b4346e990feb"

	mac := hmac.New(sha256.New, []byte(secret))
	timestamp := "1612424369000"
	msg := fmt.Sprintf("%s\n%s", timestamp, secret)
	mac.Write([]byte(msg))
	expectedMAC := mac.Sum(nil)
	expect := base64.StdEncoding.EncodeToString(expectedMAC)

	a := url.QueryEscape(expect)
	log.Println("===")
	log.Println(a)
	log.Println("===")
}
