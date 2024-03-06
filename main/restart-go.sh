#!/bin/bash
export MYSQL_MASTER="root:qwer1234@tcp(192.168.81.228:3306)/dpc_tools?charset=utf8"
export REDIS_HOST=192.168.81.228
export REDIS_PORT=6379
export REDIS_PWD=qwer1234
app_name=main
PID=$(ps -ef | grep $app_name | grep -v grep | awk '{ print $2 }')
if [ -z "$PID" ]
then
	echo  $app_name is already stopped
else
	echo kill $PID
	kill -9 $PID
        sleep 5s
	echo ' $app_name 停止服务完成'
fi

nohup go run ./main.go &
#tail -f nohup.out
