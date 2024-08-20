# gocolly-beike

### 项目描述


### 使用说明
~~~

~~~

### 代码

```bash
// 编译
go build main.go
// 调用|抓取[1]北京
main -cityid=1
```

可以写个计划任务脚本

```bash
#!/bin/bash

echo "[`date +"%Y-%m-%d %H:%M:%S"`]//////////START//////////"
city_id=(1 2 3 4 5 101)
for v in ${city_id[@]}; do
    echo "[`date +"%Y-%m-%d %H:%M:%S"`]//////////[#${v}#]//////////"
    /data/golang/project/gocolly/bin/main -cityid=$v >> /data/cronlog/beike-$v.log 2>&1
done
echo -e "[`date +"%Y-%m-%d %H:%M:%S"`]///////////END///////////\n"
```
