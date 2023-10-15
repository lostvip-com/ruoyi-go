#!/bin/sh
export GOPROXY=https://goproxy.cn,direct
CGO_ENABLED=1  go build  -ldflags "-w -s" -o  main
#构建镜像.

docker build -t reg.lostvip.com/sys:1.0 .
#发布到swarm,命名空间g
docker stack deploy -c deploy-swarm-ry.yml g
