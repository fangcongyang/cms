package sys

import (
	"cms/api/v1/sys"
	"cms/middleware"
	"github.com/gin-gonic/gin"
)

func InitJwtRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("jwt").Use(middleware.OperationRecord())
	{
		ApiRouter.POST("jsonInBlacklist", sys.JsonInBlacklist) // jwt加入黑名单
	}
}
