package main

/*
等同于docker image ls
*/
import (
	"docker.io/go-docker"
	"fmt"

	"docker.io/go-docker/api/types"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}
}
