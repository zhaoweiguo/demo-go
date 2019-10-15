package main

import (
	"github.com/drone/drone-runtime/engine"
	"log"
)

func main() {

	source := "./github.com/drone/drone-runtime/usage/1_hello_world.json"
	// 解析json文件
	config, err := engine.ParseFile(source)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("1====================")
	log.Println(config)

	log.Println("1.1 ====================")
	log.Println(config.Files)
	log.Println(config.Files[0])
	log.Println(config.Files[0].Data)
	log.Println(config.Files[0].Metadata)

	log.Println("1.2 ====================")
	log.Println(config.Docker)

	log.Println("1.3 ====================")
	log.Println(config.Metadata)
	log.Println(config.Metadata.UID)
	log.Println(config.Metadata.Labels)

	log.Println("1.4 ====================")
	for _, step := range config.Steps {
		log.Println(step)

		log.Println("1.4.1 ====================")
		log.Println(step.RunPolicy)

		log.Println(step.Docker)
		log.Println(step.Docker.Image)
		log.Println(step.Docker.User)
		log.Println(step.Docker.Args)
		log.Println(step.Docker.Command)

		log.Println(step.Volumes)

		log.Println(step.WorkingDir)

		log.Println(step.Envs)

		log.Println(step.Secrets)

		log.Println(step.Files)

		log.Println("1.4.8 ====================")
		for _, mount := range step.Files {
			log.Println(mount)
			log.Println("1.4.8.1 ====================")
			log.Println(mount.Name)
			log.Println(mount.Mode)
			log.Println(mount.Path)
		}
	}
}
