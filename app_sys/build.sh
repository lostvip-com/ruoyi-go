#!/bin/sh
#export GOPROXY=https://goproxy.cn,direct
#GOOS=linux GOARCH=amd64
#go build  -ldflags "-w -s" -o  main
#构建镜像.

docker build -t reg.lostvip.com/sys:1.2.0 .
docker push reg.lostvip.com/sys:1.2.0 .
#发布到swarm,命名空间g
docker stack deploy -c deploy-swarm-ry.yml g
docker service ls
docker service logs -f g_sys
