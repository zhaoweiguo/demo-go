package conf

import "github.com/jinzhu/configor"

var Config Kafka

type Kafka struct {
	Verbose bool
	Version string
	Oldest  bool
	Topics  string
	Group   string
	Brokers string
}

func init() {
	configor.Load(&Config, "./conf/config.yml")
}
