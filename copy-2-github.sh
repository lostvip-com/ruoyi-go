#!/bin/bash
# 目标地址
dist=/root/opensource/ruoyi-go

git reset --hard
git pull
chmod 777 *.sh
#全部覆盖过去
cp  -r  ./*           $dist/
#删除危险信息
rm -rf $dist/main/application.yml
rm -rf $dist/main/data.db
rm -rf $dist/copy-2-github.sh

# 删除开发中的数据
rm -rf $dist/ui-vue3
rm -rf $dist/bak/ui-vue3
rm -rf $dist/main/internal/iot_dev
rm -rf $dist/main/template/chromedp
rm -rf $dist/stock
# 用本地配置 覆盖配置文件
cp -r  $dist/bak/local/application.yml      $dist/main/
cp -r  $dist/bak/sql/data.db                $dist/main/
#删除敏感配置
rm -rf $dist/bak/prod
# 查看文件
cd $dist/
ls
pwd


