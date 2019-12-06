package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Shopify/sarama"
	"github.com/zhaoweiguo/demo-go/github.com/Shopify/sarama/conf"
	"github.com/zhaoweiguo/demo-go/github.com/Shopify/sarama/kafka"
)

func main() {
	confKafka := conf.Config.Kafka
	if confKafka.Verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	version, err := sarama.ParseKafkaVersion(confKafka.Version)
	if err != nil {
		panic(err)
	}

	config := sarama.NewConfig()
	config.Version = version

	if confKafka.Oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	consumer := kafka.Consumer{
		Ready: make(chan bool, 0),
	}

	ctx := context.Background()
	client, err := sarama.NewConsumerGroup(strings.Split(confKafka.Brokers, ","), confKafka.Group, config)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			err := client.Consume(ctx, strings.Split(confKafka.Topics, ","), &consumer)
			if err != nil {
				panic(err)
			}
		}
	}()

	<-consumer.Ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	<-sigterm // Await a sigterm signal before safely closing the consumer

	err = client.Close()
	if err != nil {
		panic(err)
	}

}
