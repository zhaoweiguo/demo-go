package main

import (
	"context"
	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, _ := docker.NewEnvClient()
	containerId := "b321b961987c"
	reader, err := client.ContainerLogs(ctx, containerId, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, reader)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

}
