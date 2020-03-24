package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	"github.com/zhaoweiguo/demo-go/go.mongodb.org/mongo-driver/conf"
)

type Podcast struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Author      string             `bson:"author,omitempty"`
	Tags        []string           `bson:"tags,omitempty"`
	TimeCreated time.Time          `bson:"time_created"`
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	confMongo := conf.GetMongo()
	confHosts := confMongo.Key("HOSTS").Value()
	confDBName := confMongo.Key("DB_NAME").Value()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(confHosts))

	if err != nil {
		log.Fatal("err:", err)
	}
	defer client.Disconnect(ctx)
	database := client.Database(confDBName)

	podcastsCollection := database.Collection("podcasts")
	podcast := Podcast{
		Title:       "The Polyglot Developer",
		Author:      "Nic Raboy",
		Tags:        []string{"development", "programming", "coding"},
		TimeCreated: time.Now(),
	}
	// 插入一条数据
	insertResult, err := podcastsCollection.InsertOne(ctx, podcast)
	if err != nil {
		panic(err)
	}
	log.Println(insertResult.InsertedID)

}
