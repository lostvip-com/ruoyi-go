#!/bin/sh
export GOPROXY=https://goproxy.cn,direct \
CGO_ENABLED=1 GOOS=linux GOARCH=amd64    \
go build -ldflags "-w -s" -o main  main.go
