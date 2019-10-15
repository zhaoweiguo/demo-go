package main

import (
	"github.com/docker/distribution/reference"
	"log"
	"strings"
)

func main() {
	s := "alpine:3.6"
	log.Println(refer(s)) // docker.io/library/alpine:3.6 docker.io false

	s = "registry.cn-beijing.aliyuncs.com/zhaoweiguo/java:1.0"
	log.Println(refer(s)) // registry.cn-beijing.aliyuncs.com/api_siot_ruler/java_environment:1.0 registry.cn-beijing.aliyuncs.com false

	s = "registry.cn-beijing.aliyuncs.com/zhaoweiguo/java:latest"
	log.Println(refer(s)) // registry.cn-beijing.aliyuncs.com/zhaoweiguo/java:latest registry.cn-beijing.aliyuncs.com true

	s = "alpine"
	log.Println(refer(s)) // docker.io/library/alpine:latest docker.io true

}

func refer(s string) (canonical, domain string, latest bool, err error) {
	named, err := reference.ParseNormalizedNamed(s)
	if err != nil {
		return
	}
	named = reference.TagNameOnly(named)
	return named.String(),
		reference.Domain(named),
		strings.HasSuffix(named.String(), ":latest"),
		nil

}
