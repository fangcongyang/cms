package sys

import (
	"cms/api/v1/sys"
	"cms/middleware"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api").Use(middleware.OperationRecord())
	{
		ApiRouter.POST("createApi", sys.CreateApi)               // 创建Api
		ApiRouter.POST("deleteApi", sys.DeleteApi)               // 删除Api
		ApiRouter.POST("getApiList", sys.GetApiList)             // 获取Api列表
		ApiRouter.POST("getApiById", sys.GetApiById)             // 获取单条Api消息
		ApiRouter.POST("updateApi", sys.UpdateApi)               // 更新api
		ApiRouter.POST("getAllApis", sys.GetAllApis)             // 获取所有api
		ApiRouter.DELETE("deleteApisByIds", sys.DeleteApisByIds) // 删除选中api
	}
}
