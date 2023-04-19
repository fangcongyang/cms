package sys

import (
	"cms/api/v1/sys"
	"cms/middleware"
	"github.com/gin-gonic/gin"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
	CasbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	{
		CasbinRouter.POST("updateCasbin", sys.UpdateCasbin)
		CasbinRouter.POST("getPolicyPathByAuthorityId", sys.GetPolicyPathByAuthorityId)
	}
}
