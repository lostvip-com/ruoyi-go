

### 微服务说明
简单、简洁 用最少的代码实现业务功能,去TMD三层架构，如果你愿意，你可以在controller层写sql ！
可独立运行，也可做为 k8s 或 springcloud 微服务的业务组件；不包含token或签名校验功能，交给gateway或ingress专业组件处理！

* xorm + gin
* 没有dao层
* 没有interface接口
* 支持nacos服务注册
* 兼容springcloud服务



#### 二、目录结构  
本着简单易用的原则，简化工程结构，没有dao层、没有过多的业务模块，以微服务做为业务处理单元。

~~~
main     
├── 1-------doc      // 资料备份
├── biz                                 // 业务模块
│       └── controller                  // 路由、控制器层
│       └── model                       // 模型层，支持自动建表
│       └── service                     // 业务逻辑层
│       └── vo                          // value object
├── common          // 通用模块
│       └── dto                         // 接收参数用
│       └── nacos                       // 微服务注册相关
│       └── utils                       // 工具类
├── config          // 配置处理
│       └── dto                         // 接收参数用
│       └── nacos                       // 微服务注册相关
│       └── utils                       // 工具类
├── doc                                 // swagger文档
├── resources                           // 资源文件
│       └── pages                       // gin模板引擎
│       └── static                      // 静态文件
│       └── bootstrap.yml               // 配置文件
│       └── local.yml                   // 本地配置文件，可被nacos中的覆盖
├── build-linux.bat                     // 编译脚本
~~~

#### 三、编译运行

安装swagger文档生成工具，及运行
~~~
go get -u github.com/swaggo/swag/cmd/swag

go install github.com/swaggo/swag/cmd/swag

go  mod init lostvip.com

go mod tidy

go run main.go
~~~

热加载启动： 工具 （推荐）：https://github.com/gravityblast/fresh
