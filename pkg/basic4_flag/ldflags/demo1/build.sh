#!/bin/bash

echo "default:====================="
go build
./demo1
echo "set flag:===================="
go build -ldflags="-X 'main.Version=v1.0.0'"
./demo1