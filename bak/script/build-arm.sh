source /etc/profile
xgo -ldflags "-w -s"  -out svc-$(date '+%Y%m%d_%H%M%S') --targets=linux/arm-7 .