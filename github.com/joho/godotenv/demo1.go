package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// 用法:
// $> go build demo1.go
// $> ./demo1
// 2019/09/19 11:57:31 YOURS3BUCKET YOURSECRETKEYGOESHERE
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	// now do something with s3 or whatever
	log.Println(s3Bucket, secretKey)
}
