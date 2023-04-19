package sys

import (
	"cms/api/v1/sys"
	"cms/middleware"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("").Use(middleware.NeedInit())
	{
		BaseRouter.POST("login", sys.Login)
		BaseRouter.POST("captcha", sys.Captcha)
	}
	return BaseRouter
}
