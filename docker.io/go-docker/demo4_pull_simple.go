package main

/*
Pull an image simple

*/
import (
	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"io"
	"os"

	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	out, err := cli.ImagePull(ctx, "alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()
	io.Copy(os.Stdout, out)
}
