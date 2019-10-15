package main

import (
	"context"
	"flag"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
	"time"
)

func main()  {

	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	var (
		defaultStart = yesterday.Format("20060102")
		defaultEnd   = today.Format("20060102")

		startStr = flag.String("start", defaultStart, "开始日期")
		endStr   = flag.String("end", defaultEnd, "结束日期")
	)
	start, _ := time.Parse("20060102", *startStr)
	end, _ := time.Parse("20060102", *endStr)


	mongoUri := "mongodb://root:UWyogH5VmaTZVZQJRn@s-2zed37b6f87d4dd4.mongodb.rds.aliyuncs.com:3717"
	db := "octopus_api"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("err:", err)
	}


	table1 := "octopus_gadget_info"
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	collection1 := client.Database(db).Collection(table1)

	filter := bson.D{{"time_created", bson.D{{"$gte", start}, {"$lt", end}},}}

	cur, err := collection1.Find(ctx, filter)
	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		result := struct{
			User_id string
			Hub_id string
			Gadget_type_id string
			Time_created time.Time
		}{}
		elem := &result

		//elem := &bson.D{}    // 这个也行

		//elem := &bson.M{}    // 这个也行
		err := cur.Decode(elem)
		if err != nil { log.Fatal(err) }

		log.Println(elem)
	}


}
