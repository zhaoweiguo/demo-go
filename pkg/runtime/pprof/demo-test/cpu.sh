#!/bin/zsh

go build -o demo .
./demo --cpuprofile=.cpu.prof
go tool pprof demo .cpu.prof