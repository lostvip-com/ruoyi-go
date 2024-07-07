#!/bin/bash
######################################################
# 构建镜像并推送到docker 仓库
# 构建镜像并推送到docker 仓库
######################################################
export VER_APP=1.8.0
export GOPROXY=https://goproxy.cn,direct
go build -ldflags "-w -s" -o main  main.go
echo "------ GO 编译完成 docker 构建开始 ---------------------------"
docker images
echo "服务升级 ======>"
#read -p "输入版本号:"  ver
#export VER_APP=$ver
echo "您输入的版本号是：$VER_APP"
read -p "请选Y/N:"  para
case $para in
	[yY])
		echo " 开始升级....."
		;;
	[nN])
		echo " 即将退出！！！"
		exit 1
esac # end case
SVC_NAME=g_sys
IMG_NAME=reg.lostvip.com/lv-sys:$VER_APP
echo "------ $image 构建开始 ---------------------------"
docker build -t  $IMG_NAME .
docker push  $IMG_NAME
echo "------ $IMG_NAME 构建完成 ---------------------------"

#docker stack deploy -c ../docs/swarm/deploy-swarm-ry.yml g
#docker service ls
#docker service logs -f $SVC_NAME --tail=100
