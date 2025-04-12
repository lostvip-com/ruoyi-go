## grpc驱动

### 安装编译工具 

* 安装protoc-gen-go和protoc-gen-go-grpc
~~~
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
~~~

* 下载 protobuf 工具

> https://github.com/protocolbuffers/protobuf/releases
> 
### 生成代码 

~~~
make
~~~
