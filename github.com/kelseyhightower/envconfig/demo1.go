package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type (

	// Config provides the system configuration.
	Config struct {
		Docker     Docker
		Logging    Logging
		Registries Registries
	}

	// Docker provides docker configuration
	Docker struct {
		Config string `envconfig:"DRONE_DOCKER_CONFIG"`
	}

	// Logging provides the logging configuration.
	Logging struct {
		Debug bool `envconfig:"DRONE_LOGS_DEBUG"`
		Trace bool `envconfig:"DRONE_LOGS_TRACE"`
	}

	// Registries provides the registry configuration.
	Registries struct {
		Endpoint   string `envconfig:"DRONE_REGISTRY_ENDPOINT"`
		Password   string `envconfig:"DRONE_REGISTRY_SECRET" default:"passwd"`
		SkipVerify bool   `envconfig:"DRONE_REGISTRY_SKIP_VERIFY"`
	}
)

// $> go build demo1.go
// $> ./demo1
// 2019/09/19 12:13:20 {{} {false false} { passwd false}}
// $> DRONE_DOCKER_CONFIG=abc ./demo1
// 2019/09/19 12:15:45 {{abc} {false false} { passwd false}}
// $> DRONE_DOCKER_CONFIG=abc DRONE_LOGS_TRACE=true DRONE_REGISTRY_SECRET=secret  ./demo1
// 2019/09/19 12:17:00 {{abc} {false true} { secret false}}
func main() {

	cfg := Config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Println(err)
	}
	log.Println(cfg)

}
