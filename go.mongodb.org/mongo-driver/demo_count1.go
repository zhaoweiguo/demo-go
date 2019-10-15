package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {

	log.Println("start...")
	mongoUri := "mongodb://10.128.132.36:27017"
	//mongoUri := "mongodb://user:passwd@s-xxxxx.mongodb.rds.aliyuncs.com:3717"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal("err:", err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	collection := client.Database("octopus").Collection("octopus_gadget_info")

	filter := bson.D{}
	count, err := collection.CountDocuments(ctx, filter)
	log.Println(count)
}



