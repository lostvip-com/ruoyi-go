<p align="center">
	<img alt="logo" src="https://oscimg.oschina.net/oscnet/up-dd77653d7c9f197dd9d93684f3c8dcfbab6.png">
</p>
<h1 align="center" style="margin: 30px 0 30px; font-weight: bold;"> lv-framework v1.1.4 </h1>
<h4 align="center">Go语言版轻量级快速开发框架</h4>


## 框架简介
摒弃过度封装，代码风格极尽简洁，适合中小项目使用！<br>
架构思路沿袭着若依的以辅助生成重复代码为主，不过度封装，生成的代码可以快速修改适应不同的需求。
* 项目结构模仿springboot，简单改造后可注册到nacos中做为springcloud微服务的一员，对java开发人员友好。
* 支持类mybatis查询，sql与go代码分离（基于gorm + golang的template语法实现,为了idea能识别sql关键字，后缀使用_mapper.sql）。
* 支持在yaml文件中使用表达式从环境变量中取值，便于切换生产、开发、测试环境参数（模仿springboot）,如：
~~~
    host: ${REDIS_HOST:lostvip.com}
    如: 环境变量添加 REDIS_HOST=192.168.88.114;REDIS_PORT=6379;REDIS_PWD=dpctest
~~~

## 核心技术及组件

* DAO层架: gorm,简单易用,支持namedSQL占位符查询.</br>
* Web服务框架: Gin 简单高效，模板可热加载.</br>
* 前端技术: 模板引擎，服务端渲染.</br>
* 缓存支持: redis .</br>

## 内置功能

1.  缓存，支持redis/RAM 缓存
2.  数据库，支持mysql，sqlite3，可自行扩展其他数据库。
3.  Web服务，支持restful风格，支持跨域，支持自定义模板引擎，支持服务端渲染。：配置系统用户所属担任职务。
4.  常用工具类封装。


## 目录结构
本着简单易用的原则，简化工程结构，没有dao层。

~~~
├── lv_framework       // 依赖库，通用框架，简单封装了数据库，web等常用功能
│       └── lv_cache               // 通用缓存
│       └── lv_conf                // 通用配置
│       └── lv_db                  // 数据库相关
│       └── lv_log                 // 通用日志处理，统一风格，方便切换第三方日志库
│       └── utils                  // 通用工具类
│       └── web                    // 通用web组件  
~~~

引入
~~~
go get github.com/lostvip-com/lv-framework
~~~

## QQ交流群
如有兴趣一起完善，进QQ群 43862272

