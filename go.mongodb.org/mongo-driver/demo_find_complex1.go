package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main()  {

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

	// 条件:
	// 	 1. 日期大于xxx
	//   2. name in [1,2,3,4]
	timeString := "2019-06-02 14:34:00.998011694 +0800 CST"
	start, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timeString)
	ids := []string{"a", "b"}
	filter := bson.D{{"$and",
		bson.A{
			bson.D{{"name", bson.D{{"$in", bson.A{1,2,3,4} }} }},  // 第一种in方式
			bson.D{{"name", bson.D{{"$in", ids }} }},	// 第二种in方式
			bson.D{{"time_created", bson.D{{"$gt", start}}}},
		}}}

	//filter := bson.D{{"gadget_type_id", "9000006"}}
	//filter := bson.D{{"time_created", bson.D{{"$gt", start}}}}

	cur, err := collection.Find(ctx, filter)
	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
		// do something with result....

		log.Println(result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

