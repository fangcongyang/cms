package initialize

import (
	_ "cms/docs"
	"cms/global"
	"cms/middleware"
	"cms/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	// 配置静态文件映射
	Router.StaticFS("static", http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.GVA_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitSpiderRouter(PublicGroup)
	}

	// api
	ApiPublicGroup := Router.Group("/api")
	//ApiPublicGroup.Use(middleware.ApiAuth())
	{
		router.InitApiFictionRouter(ApiPublicGroup)
	}

	// 方便统一添加路由组前缀 多服务器上线使用
	CmsPublicGroup := Router.Group("/cms")
	{
		router.InitSysPublicRouter(CmsPublicGroup)
	}
	CmsPrivateGroup := Router.Group("/cms")
	CmsPrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		router.InitSysPrivateRouter(CmsPrivateGroup)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}
