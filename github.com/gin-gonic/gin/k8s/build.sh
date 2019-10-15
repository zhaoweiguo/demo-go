#!/usr/bin/env bash

GOARCH=amd64 GOOS=linux go build -o gordondemo ../demo1_staticfile.go

docker build -t registry.cn-beijing.aliyuncs.com/octopus-test/gordondemo:v8 .
docker push registry.cn-beijing.aliyuncs.com/octopus-test/gordondemo:v8
kubectl apply -f ./gordondemo_pod.yml



