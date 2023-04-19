package sys

import (
	"cms/api/v1/sys"
	"cms/middleware"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(Router *gin.RouterGroup) {
	SystemRouter := Router.Group("system").Use(middleware.OperationRecord())
	{
		SystemRouter.POST("getSystemConfig", sys.GetSystemConfig) // 获取配置文件内容
		SystemRouter.POST("setSystemConfig", sys.SetSystemConfig) // 设置配置文件内容
		SystemRouter.POST("getServerInfo", sys.GetServerInfo)     // 获取服务器信息
		SystemRouter.POST("reloadSystem", sys.ReloadSystem)       // 重启服务
	}
}
