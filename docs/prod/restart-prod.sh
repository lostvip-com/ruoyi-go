#!/bin/bash
export MYSQL_MASTER="root:Ssz123456!@tcp(lostvip.com:3307)/test?charset=utf8"
export REDIS_HOST=lostvip.com
export REDIS_PORT=6379
export REDIS_PWD=Ssz123456!
app_name=main
PID=$(ps -ef | grep $app_name | grep -v grep | awk '{ print $2 }')
if [ -z "$PID" ]
then
	echo  $app_name is already stopped !!
else
	echo kill $PID
	kill -9 $PID
        sleep 5s
	echo  $app_name 停止服务完成!
fi

nohup  ./main &
ps -ef|grep main

tail -f nohup.out 
