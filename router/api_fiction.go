package router

import (
	v1 "cms/api/v1"
	"github.com/gin-gonic/gin"
)

//app相关路由

func InitApiFictionRouter(Router *gin.RouterGroup) {
	//小说分类路由
	FictionSortRouter := Router.Group("fictionSort")
	{
		FictionSortRouter.GET("", v1.QueryFictionSortPageList)
	}

	//小说路由
	FictionRouter := Router.Group("fiction")
	{
		FictionRouter.GET("", v1.QueryFictionPageList)
		FictionRouter.GET(":id", v1.GetFictionById)
	}

	//小说章节路由
	FictionChapterRouter := Router.Group("fictionChapter")
	{
		FictionChapterRouter.GET("", v1.QueryFictionChapterList)
		FictionChapterRouter.GET(":id", v1.GetFictionChapterById)
	}

	//登录路由
	LoginRouter := Router.Group("login")
	{
		LoginRouter.POST("", v1.ApiLogin)
	}

	//导航路由
	BannerRouter := Router.Group("base")
	{
		BannerRouter.POST("banner/queryList", v1.ApiBanner)
	}

	//用户小说路由
	UserFictionRouter := Router.Group("userFiction")
	{
		UserFictionRouter.GET("", v1.ListBookShelf)
		UserFictionRouter.POST("save", v1.AddBookShelf)
	}
}
