## 框架简介
摒弃过度封装，代码风格极尽简洁，适合中小项目使用！<br>
架构思路沿袭着若依的以辅助生成重复代码为主，不过度封装，生成的代码可以快速修改适应不同的需求。
* 项目结构模仿springboot，简单改造后可注册到nacos中做为springcloud微服务的一员，对java开发人员友好。
* 支持类mybatis查询，sql与go代码分离（基于gorm + golang的template语法实现,为了idea能识别sql关键字，后缀使用_mapper.sql）。
* 提供相关服务配置脚本，可发布到 k8s、docker swarm中做为微服务使用。
* 支持在yaml文件中使用表达式从环境变量中取值，便于切换生产、开发、测试环境参数（模仿springboot）,如：
~~~
    host: ${REDIS_HOST:lostvip.com}
    如: 环境变量添加 REDIS_HOST=192.168.88.114;REDIS_PORT=6379;REDIS_PWD=dpctest
~~~
* 1.x版本是基于xorm开发，2.0之后的版本将全面切换到gorm。
## 核心技术及组件

* DAO层架: gorm,简单易用,支持namedSQL占位符查询.</br>
* Web服务框架: Gin 简单高效，模板可热加载.</br>
* 前端技术: 模板引擎，服务端渲染.</br>
* 缓存支持: redis .</br>
* 安全认证：支持通过jwt无状态 或 token有状态的安全校验.
* 代码生成：在线生成代码，并一键导入菜单到业务模块.
* 导出excel文件  tealeg/xlsx.
* api文档生成    swaggo/swag.      
* 图形验证码     base64Captcha.    
* 服务器监控     gopsutil.   
* 若依前端组件   RuoYi  v4.7.0 .

## 内置功能

1.  用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2.  部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3.  岗位管理：配置系统用户所属担任职务。
4.  菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5.  角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6.  字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7.  参数管理：对系统动态配置常用参数。
8.  通知公告：系统通知公告信息发布维护。
9.  操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
10.  登录日志：系统登录日志记录查询包含登录异常。
11.  在线用户：当前系统中活跃用户状态监控。
12.  定时任务：在线（添加、修改、删除)任务调度包含执行结果日志。
13.  代码生成：前后端代码的生成（Go、html、json、sql） 。
14.  系统接口：根据业务代码自动生成相关的api接口文档。
15.  服务监控：监视当前系统CPU、内存、磁盘、堆栈等相关信息。
16.  在线构建器：拖动表单元素生成相应的HTML代码。
17.  案例演示：常用的前端组件整合演示。


## 目录结构
本着简单易用的原则，简化工程结构，没有dao层。

~~~
main[]()     
├── docs               // 资料备份
│   └── deploy-swarm-ry.yml      // docker swarm 服务编排脚本
├── main               // 系统管理
│   └── internal                 // 源码存放目录,internal禁止其它模块引用
│   └── static                   // 静态文件目录
│   └── mapper                   // 类mybatis映射文件存放目录
│   └── template                 // 模板存放目录
│   └── bootstrap.yml            // 项目根配置文件，模仿springboot
│   └── application.ym           // 项目配置文件，模仿springboot
│   └── build.sh                 // 编译脚本，方便打包
│   └── build-cgo.sh             // 编译脚本，启用cgo,需要打包sqlite时有用
│   └── build-linux.bat          // 编译脚本，windows环境交叉编译用
│   └── data.db                  // 使用sqlite数据库时有用，只使用mysql可删除                             
│   └── Dockerfile               // Docker镜像文件
├── lv_framework       // 框架通用模块
│       └── cache               // 通用缓存处理
│       └── conf                // 通用配置处理
│       └── db                  // 数据库配置处理
│       └── utils               // 通用工具类
│       └── logme               // 通用日志处理，统一风格，方便切换第三方日志库
│       └── web                 // 通用web组件
│       
├── go.work            // 项目模块配置
├── build-linux.bat    // 编译脚本
~~~

## 启动
前置条件：
> 1. 配置GO_PROXY环境变量，如：export GO_PROXY=https://goproxy.cn
安装相关依赖！
### 1. 修改应该的数据库及缓存配置
* sqlite：默认使用sqlite数据库，可不做任何修改
* mysql： 导入docs中的sql文件，修改application.yml中的mysql相关信息

### 2. cd 进入main目录启动应用
* 方式一 ：常规启动
~~~
go run main.go
~~~
启动完成后，在控制台会打印出访问地址： http://127.0.0.1:8080
* 方式二： 使用第三方热加载组件
> 前置条件： cd 到 [main](main) 目录，安装热加载框架：
> <br>  go get github.com/pilu/fresh
> <br>  go install github.com/pilu/fresh

启动
~~~
fresh
~~~
## 登录账号和密码
> 演示1: http://demo.lostvip.com <p> 
> 账号: admin / admin123  <p>
>
注意,本站未做任何防护,管理员有所有权限,请不要删除任何数据!


## 演示图

<img alt="" height="400" src="bak/1.jpg" width="600" />
<img alt="" height="400" src="bak/2.jpg" width="600" />
<img alt="" height="400" src="bak/3.png" width="600" />
<img alt="" height="400" src="bak/4.png" width="600" />

## 感谢(排名不分先后)
> xorm  https://xorm.io/zh/docs/   
> 
> web框架 [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
>
> 缓存框架  https://github.com/redis/go-redis
>
> 配置文档   [https://github.com/BurntSushi/toml](https://github.com/BurntSushi/toml)
>
> RuoYi框架 [https://github.com/yangzongzhuan/RuoYi](https://github.com/yangzongzhuan/RuoYi)
>
> tealeg [https://github.com/tealeg/xlsx](https://github.com/tealeg/xlsx)
>
> swaggo [https://github.com/swaggo/swag](https://github.com/swaggo/swag)
>
> 若依goframe版 https://github.com/guolingege/yjgo
>
> 热加载框架  http://github.com/pilu/fresh
>
> sql文件解析: https://github.com/qustavo/dotsql
## 项目支持


## QQ交流群

如有兴趣一起完善，进QQ群 43862272

<img alt="" height="389" src="https://github.com/lostvip-com/ruoyi-go/blob/main/docs/qq.png" width="300" />

 <p>Gitee <a href="https://gitee.com/lostvip_com/ruoyi-go">https://gitee.com/lostvip_com/ruoyi-go</a></p>
 <p>Github <a href="https://github.com/lostvip-com/ruoyi-go">https://github.com/lostvip-com/ruoyi-go</a></p>
 <p>Gitcode <a href="https://gitcode.com/lostvip-com1/ruoyi-go/tree/main">https://gitcode.com/lostvip-com1/ruoyi-go/tree/main</a></p>
