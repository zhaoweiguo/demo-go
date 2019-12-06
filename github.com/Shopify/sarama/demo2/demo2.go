package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	validVersions := []string{"0.8.2.0", "0.8.2.1", "0.9.0.0", "0.10.2.0", "1.0.0", "2.1.1", "2.2.0"}
	for _, s := range validVersions {
		v, err := sarama.ParseKafkaVersion(s)
		if err != nil {
			log.Printf("could not parse valid version %s: %s", s, err)
		}
		if v.String() != s {
			log.Printf("version %s != %s", v.String(), s)
		}
		log.Println(v)
	}

	log.Println("---------------")

	invalidVersions := []string{"0.8.2-4", "0.8.20", "1.19.0.0", "1.0.x"}
	for _, s := range invalidVersions {
		_, err := sarama.ParseKafkaVersion(s)
		if err == nil {
			log.Printf("invalid version %s parsed without error", s)
		}
		log.Println(err)
	}

}
