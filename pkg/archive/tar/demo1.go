package main

import (
	"archive/tar"
	"bytes"
	"log"
	"time"
)

func main() {
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
	log.Println(w.Bytes())

}
