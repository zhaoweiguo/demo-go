package main

import (
	"crypto/sha256"
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {

	secret := "SECc4e8395e49471c7260be9057b9a2926b0082221ddb578786752bebb46f183a6a"

	timestamp := "1612420355000"
	msg := fmt.Sprintf("%s\n%s", timestamp, secret)

	h := sha256.New()
	h.Write([]byte(msg))

	log.Printf("%x", h.Sum(nil))

}
