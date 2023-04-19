package router

import (
	v1 "cms/api/v1"
	"github.com/gin-gonic/gin"
)

func InitSpiderRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("spider")
	{
		ApiRouter.GET("startFictionSpider", v1.StartFictionSpider)               // 创建Api
		ApiRouter.GET("saveFictionToDataBase", v1.SaveFictionToDataBase)
	}
}

