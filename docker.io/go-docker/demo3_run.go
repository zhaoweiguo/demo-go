package main

/*
from: https://docs.docker.com/develop/sdk/examples/
title: Run a container
等同于: docker run alpine echo hello world
*/
import (
	"context"
	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
	"log"

	"io"
	"os"
)

func main() {
	ctx := context.Background()
	cli, _ := docker.NewEnvClient()

	// 等同于docker pull
	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	// 等同于 docker run -it alpine echo "hello world"
	// 等同于docker create...
	config := &container.Config{
		Image: "alpine",
		Cmd:   []string{"sh"},
		Tty:   true,
	}
	resp, err := cli.ContainerCreate(ctx, config, nil, nil, "")
	if err != nil {
		panic(err)
	}

	log.Println(resp)
	log.Println(resp.ID, resp.Warnings)

	// 等同于docker start
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	// 等待容器运行完成
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	// 等同于docker logs
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

}
