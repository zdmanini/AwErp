package main

import (
	adminRouters "Awesome/admin/routers"
	admin "Awesome/admin/service"
	"Awesome/config"
	"Awesome/core"
	"Awesome/core/response"
	"Awesome/dongming"
	genRouters "Awesome/generator/routers"
	gen "Awesome/generator/service"
	"Awesome/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// initDI 初始化DI
func initDI() {
	regFunctions := admin.InitFunctions
	regFunctions = append(regFunctions, gen.InitFunctions...)
	regFunctions = append(regFunctions, core.GetDB)
	for i := 0; i < len(regFunctions); i++ {
		if err := core.ProvideForDI(regFunctions[i]); err != nil {
			log.Fatalln(err)
		}
	}
}

// initRouter 初始化router
func initRouter() *gin.Engine {
	// 初始化gin
	gin.SetMode(config.Config.GinMode)
	router := gin.New()
	// 设置静态路径
	router.Static(config.Config.PublicPrefix, config.Config.UploadDirectory)
	router.Static(config.Config.StaticPath, config.Config.StaticDirectory)
	// 设置中间件
	router.Use(gin.Logger(), middleware.Cors(), middleware.ErrorRecover())
	// 演示模式
	if config.Config.DisallowModify {
		router.Use(middleware.ShowMode())
	}
	// 特殊异常处理
	router.NoMethod(response.NoMethod)
	router.NoRoute(response.NoRoute)
	// 注册路由
	group := router.Group("/admin")

	routers := adminRouters.InitRouters[:]
	routers = append(routers, genRouters.InitRouters...)
	for i := 0; i < len(routers); i++ {
		core.RegisterGroup(group, routers[i])
	}
	return router
}

// initServer 初始化server
func initServer(router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           ":" + strconv.Itoa(config.Config.ServerPort),
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func main() {
	// 刷新日志缓冲
	defer core.Logger.Sync()
	// 程序结束前关闭数据库连接
	if core.GetDB() != nil {
		db, _ := core.GetDB().DB()
		defer db.Close()
	}
	// 初始化配置
	go dongming.Run()
	// 初始化DI
	initDI()
	// 初始化router
	router := initRouter()
	// 初始化server
	s := initServer(router)
	// 运行服务
	log.Fatalln(s.ListenAndServe().Error())
}
