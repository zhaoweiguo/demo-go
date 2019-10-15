package main

import (
	"archive/tar"
	"bytes"
	"context"
	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	copyOpts := types.CopyToContainerOptions{}
	copyOpts.AllowOverwriteDirWithFile = false

	container := "9a6715c11605" // 容器id
	var tar = tarfile()
	err = cli.CopyToContainer(ctx, container, "/", bytes.NewReader(tar), copyOpts)

}

func tarfile() []byte {

	path := "/tmp/abc"
	mode := int64(511)
	data := []byte("ZWNobyBoZWxsbyB3b3JsZAo=")

	w := new(bytes.Buffer)
	t := tar.NewWriter(w)
	h := &tar.Header{
		Name:    path,
		Mode:    mode,
		Size:    int64(len(data)),
		ModTime: time.Now(),
	}

	err := t.WriteHeader(h)
	log.Println(err)
	num, err := t.Write(data)
	log.Println(num, err)
	return w.Bytes()
}
