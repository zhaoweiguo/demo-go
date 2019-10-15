#!/usr/bin/env bash

GOARCH=amd64 GOOS=linux go build -o demo6 ../demo6_client_https.go

docker build -t registry.cn-hangzhou.aliyuncs.com/xinxi/test:gitcode1 .
#docker push registry.cn-hangzhou.aliyuncs.com/xinxi/test:v2
