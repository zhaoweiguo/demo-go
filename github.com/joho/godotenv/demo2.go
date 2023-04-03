package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// 用法:
// $> go build demo2_user.go
// $> ./demo2_user
// 2019/09/19 12:03:27
// 2019/09/19 11:59:32 another_YOURS3BUCKET another_YOURSECRETKEYGOESHERE
// $> S3_BUCKET=help ./demo2_user
// 2019/09/19 12:02:35 help
// 2019/09/19 12:00:21 help another_YOURSECRETKEYGOESHERE
// $> S3_BUCKET=help SECRET_KEY=hello ./demo2_user
// 2019/09/19 12:01:18 help hello
// 2019/09/19 12:01:18 help hello
func main() {
	s3Bucket0 := os.Getenv("S3_BUCKET")
	secretKey0 := os.Getenv("SECRET_KEY")
	log.Println(s3Bucket0, secretKey0)

	err := godotenv.Load(".anotherEnv")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	// now do something with s3 or whatever
	log.Println(s3Bucket, secretKey)
}
