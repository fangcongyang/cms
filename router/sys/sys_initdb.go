package sys

import (
	"cms/api/v1/sys"
	"github.com/gin-gonic/gin"
)

func InitInitRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("init")
	{
		ApiRouter.POST("initdb", sys.InitDB)   // 创建Api
		ApiRouter.POST("checkdb", sys.CheckDB) // 创建Api
	}
}
