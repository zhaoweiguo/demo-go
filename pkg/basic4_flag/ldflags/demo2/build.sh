#!/bin/bash

echo "default:====================="
go build
./demo2
echo "set flag:===================="
go build -ldflags "-X 'github.com/zhaoweiguo/demo-go/pkg/basic4_flag/ldflags/demo2/app/config.Version=0.0.1' -X 'github.com/zhaoweiguo/demo-go/pkg/basic4_flag/ldflags/demo2/app/config.BuildTime=$(date)'"
./demo2

go tool nm ./demo2 | grep demo2 | less