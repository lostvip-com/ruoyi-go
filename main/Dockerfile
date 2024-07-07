FROM ubuntu

LABEL maintainer="lostvip.com"
###############################################################################
#                                INSTALLATION
###############################################################################
RUN apt-get update -y && apt-get upgrade -y
RUN apt-get install -y tzdata && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN apt-get install -y inetutils-ping
RUN apt-get install -y telnet
RUN apt-get install -y wget

# 设置固定的项目路径
ENV WORKDIR /root
# 添加I18N多语言文件、静态文件、配置文件、模板文件
ADD main                 $WORKDIR/main
ADD data.db              $WORKDIR/data.db
ADD application.yml      $WORKDIR/application.yml
ADD bootstrap.yml        $WORKDIR/bootstrap.yml
ADD static               $WORKDIR/static/
ADD template             $WORKDIR/template/
###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
RUN chmod +x $WORKDIR/main
RUN pwd           # 查看当前工作目录
RUN ls -al        # 查看工作目录的内容
EXPOSE 3001 3001
CMD ["/root/main"]



