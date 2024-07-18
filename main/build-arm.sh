source /etc/profile
export GO_PROXY=https://goproxy.cn
xgo -ldflags "-w -s"  -out app-0.1 --targets=linux/arm-7 .
#xgo  -out app-0.1 --targets=linux/arm-7 .
