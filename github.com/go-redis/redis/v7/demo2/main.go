package main

import (
	"github.com/go-redis/redis/v7"
	"log"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     Config.Url,
		Password: Config.Pwd, // "" no password set
		DB:       0,          // use default DB
	})

	pong, err := client.Ping().Result()
	log.Println(pong, err)

	stringSliceCmds := client.Keys("*")
	for _, item := range stringSliceCmds.Val() {
		log.Println(item)
	}

}
