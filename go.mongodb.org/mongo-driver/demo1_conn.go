package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main()  {
	normalConn()
	userPasswdConn()
}


func normalConn() {
	log.Println("start...")
	mongoUri := "mongodb://10.128.132.36:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal("err:", err)
	}
}

func userPasswdConn()  {
	log.Println("start...")
	mongoUri := "mongodb://user:password@xxx.mongodb.rds.aliyuncs.com:3717"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal("err:", err)
	}
}





