package server

import (
	"fmt"
	"github.com/spf13/cast"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"html/template"
	"lostvip.com/conf"
	"lostvip.com/logme"
	"lostvip.com/web/gintemplate"
	"lostvip.com/web/middleware"
	"lostvip.com/web/router"
	"robvi/app/common/global"
	"robvi/app/system/service/system/calcu"
	"robvi/app/system/service/system/config"
	"robvi/app/system/service/system/dict"
	"robvi/app/system/service/system/permission"

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
	server := &http.Server{
		Addr:           mySvr.Addr,
		Handler:        mySvr.Handler,
		ReadTimeout:    mySvr.ReadTimeout,
		WriteTimeout:   mySvr.WriteTimeout,
		MaxHeaderBytes: mySvr.MaxHeaderBytes,
	}

	log.Printf("[%v]Server listen: %v Actual pid is %d", mySvr.ServerName, mySvr.Addr, syscall.Getpid())

	println("app服务启动成功！！")
	host := conf.Config().GetServerIP()
	path := conf.Config().GetContextPath()
	port := cast.ToString(conf.Config().GetServerPort())
	fmt.Println("##############################################################")
	fmt.Println("go.application.name: " + global.GetConfigInstance().GetAppName())
	fmt.Println("go.datasource.master: " + global.GetConfigInstance().GetMaster())
	//加载模板引擎
	fmt.Println("http://" + host + ":" + port + strings.ReplaceAll(path, "//", "/"))
	fmt.Println("http://127.0.0.l:" + port + strings.ReplaceAll(path+"/swagger/index.html", "//", "/"))
	fmt.Println("http://" + host + ":" + port + strings.ReplaceAll(path+"/swagger/index.html", "//", "/"))
	fmt.Println("##############################################################")
	server.ListenAndServe()
}

// 创建服务
func New(addr string) *MyServer {
	gin.DefaultWriter = logme.Log.Logger.Out
	contextPath := conf.Config().GetContextPath()
	var s MyServer
	s.WriteTimeout = 60 * time.Second
	s.ReadTimeout = 60 * time.Second
	s.Addr = addr
	s.ServerName = conf.Config().GetAppName()
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
	//////////////////////////////////////////////////////////////////////////////////
	// web 页面
	/////////////////////////////////////////////////////////////////////////////////
	routerBase := engine.Group(contextPath)
	routerBase.GET("/swagger/*any", gs.DisablingWrapHandler(swaggerFiles.Handler, conf.KEY_SWAGGER_OFF))
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
		Partials:  []string{"header", "footer", "system/menu/icon"},
		Funcs: template.FuncMap{
			"hasPermi":          permission.HasPermi,
			"getPermiButton":    permission.GetPermiButton,
			"getDictLabel":      dict.GetDictLabel,
			"getDictTypeSelect": dict.GetDictTypeSelect,
			"getDictTypeRadio":  dict.GetDictTypeRadio,
			"getDictTypeData":   dict.GetDictTypeData,
			"Copyright":         config.GetCopyright,
			"OssUrl":            config.GetOssUrl,
			"Ctx":               config.GetCtx,
			"getCtxPath":        config.GetCtxPath,
			"addInt":            calcu.AddInt,
		},
		DisableCache: true,
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
