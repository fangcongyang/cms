package sys

import (
	"cms/api/v1/sys"
	"github.com/gin-gonic/gin"
)

func InitAutoCodeRouter(Router *gin.RouterGroup) {
	AutoCodeRouter := Router.Group("autoCode")
	{
		AutoCodeRouter.POST("preview", sys.PreviewTemp)   // 获取自动创建代码预览
		AutoCodeRouter.POST("createTemp", sys.CreateTemp) // 创建自动化代码
		AutoCodeRouter.GET("getTables", sys.GetTables)    // 获取对应数据库的表
		AutoCodeRouter.GET("getDB", sys.GetDB)            // 获取数据库
		AutoCodeRouter.GET("getColumn", sys.GetColumn)    // 获取指定表所有字段信息
	}
}
