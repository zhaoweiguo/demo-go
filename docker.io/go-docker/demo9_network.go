package main

import (
	"context"
	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"log"
)

func main() {

	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	driver := "bridge"
	labels := map[string]string{}
	resp, err := cli.NetworkCreate(ctx, "uid_W6kaKXlwtVX4gV6L", types.NetworkCreate{
		Driver: driver,
		Labels: labels,
	})
	log.Println(resp, err)

}
