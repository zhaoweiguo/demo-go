#!/bin/bash

echo "default:====================="
go build
./demo2_user
echo "set flag:===================="
go build -ldflags "-X 'github.com/zhaoweiguo/demo-go/pkg/basic4_flag/ldflags/demo2/app/config.Version=0.0.1' -X 'github.com/zhaoweiguo/demo-go/pkg/basic4_flag/ldflags/demo2/app/config.BuildTime=$(date)'"
./demo2_user

go tool nm ./demo2_user | grep demo2_user | less