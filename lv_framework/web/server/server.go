package server

import (
	"context"
	"fmt"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/web/gintemplate"
	"github.com/lostvip-com/lv_framework/web/middleware"
	"github.com/lostvip-com/lv_framework/web/router"
	"github.com/spf13/cast"
	"html/template"
	//gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"
)

// HTTP服务结构体
type MyServer struct {
	server         *http.Server
	ServerName     string        //服务名称
	Addr           string        //监听地址端口
	ServerRoot     string        //静态资源文件夹
	Handler        *gin.Engine   //HTTP Handler
	ReadTimeout    time.Duration //读取超时时间
	WriteTimeout   time.Duration //写入超时时间
	MaxHeaderBytes int           //http头大小设置
}

// 启动服务
func (mySvr *MyServer) Start() {
	mySvr.server = &http.Server{
		Addr:           mySvr.Addr,
		Handler:        mySvr.Handler,
		ReadTimeout:    mySvr.ReadTimeout,
		WriteTimeout:   mySvr.WriteTimeout,
		MaxHeaderBytes: mySvr.MaxHeaderBytes,
	}

	log.Printf("[%v]Server listen: %v Actual pid is %d", mySvr.ServerName, mySvr.Addr, syscall.Getpid())
	host := lv_global.Config().GetServerIP()
	path := lv_global.Config().GetContextPath()
	port := cast.ToString(lv_global.Config().GetServerPort())
	cache := lv_global.Config().GetValueStr("go.cache")
	fmt.Println("##############################################################")
	fmt.Println("go.application.name: " + lv_global.Config().GetAppName())
	fmt.Println("go.cache: " + cache)
	if cache == "redis" {
		fmt.Println("go.redis.host: " + lv_global.Config().GetValueStr("go.redis.host"))
	}
	fmt.Println("go.datasource.master: " + lv_global.Config().GetMaster())
	//加载模板引擎
	fmt.Println("http://localhost:" + port + strings.ReplaceAll(path, "//", "/"))
	fmt.Println("http://localhost:" + port + strings.ReplaceAll(path+"/swagger/index.html", "//", "/"))
	fmt.Println("http://" + host + ":" + port + strings.ReplaceAll(path+"/swagger/index.html", "//", "/"))
	fmt.Println("##############################################################")
	mySvr.server.ListenAndServe()
}

func (mySvr *MyServer) ShutDown() {
	mySvr.server.Shutdown(context.Background())
}

// 创建服务
func New(addr string) *MyServer {
	gin.DefaultWriter = lv_log.GetLog().Out
	contextPath := lv_global.Config().GetContextPath()
	var s MyServer
	s.WriteTimeout = 60 * time.Second
	s.ReadTimeout = 60 * time.Second
	s.Addr = addr
	s.ServerName = lv_global.Config().GetAppName()
	s.MaxHeaderBytes = 1 << 20
	s.Handler = InitGinRouter(contextPath)

	return &s
}

func InitGinRouter(contextPath string) *gin.Engine {
	engine := gin.Default()
	///////////////////////中间件处理start////////////////////////////////////////////////
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{}))
	engine.Use(middleware.RecoverError)   // 全局异常处理,自定义错误处理
	engine.Use(middleware.SetTraceId)     // traceId
	engine.Use(middleware.Options)        // 跨域处理
	engine.Use(middleware.LoggerToFile()) //日志
	//engine.Use(middleware.RateLimit())  // 限流
	//router.Use(gzip.Gzip(gzip.DefaultCompression)),开启后客户端无法收到，尚未解决此问题不要打开
	//router.Use(Secure)
	engine.Use(middleware.IfProxyForward())
	//////////////////////////////////////////////////////////////////////////////////
	// web 页面
	/////////////////////////////////////////////////////////////////////////////////
	routerBase := engine.Group(contextPath)
	//routerBase.GET("/swagger/*any", gs.DisablingWrapHandler(swaggerFiles.Handler, lv_conf.KEY_SWAGGER_OFF))
	tmp, _ := os.Getwd()
	staticPath := tmp + "/static"
	fmt.Println("设置静态目录：" + staticPath)

	routerBase.StaticFS("/static", http.Dir(staticPath))
	routerBase.StaticFile("/favicon.ico", staticPath+"/favicon.ico")
	//加载模板引擎
	engine.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "template",
		Extension: ".html",
		Master:    "",
		Partials:  lv_global.Config().GetPartials(), //[]string{"header", "footer", "system/menu/icon"}
		Funcs:     template.FuncMap(lv_global.Config().GetFuncMap()),
		CacheTpl:  lv_global.Config().IsCacheTpl(),
	})
	//注册路由
	if len(router.GroupList) > 0 {
		for _, group := range router.GroupList {
			grp := routerBase.Group(group.RelativePath, group.Handlers...)
			for _, r := range group.Router {
				if r.Method == "ANY" {
					grp.Any(r.RelativePath, r.HandlerFunc...)
				} else {
					grp.Handle(r.Method, r.RelativePath, r.HandlerFunc...)
				}
			}
		}
	}
	return engine
}
