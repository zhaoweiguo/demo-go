package kafka

import (
	"github.com/Shopify/sarama"
	"log"
)

func Handle(message *sarama.ConsumerMessage) {
	log.Println(message)
}