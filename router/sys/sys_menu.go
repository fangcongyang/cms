package sys

import (
	"cms/api/v1/sys"
	"cms/middleware"
	"github.com/gin-gonic/gin"
)

func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	MenuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	{
		MenuRouter.POST("getMenu", sys.GetMenu)                 // 获取菜单树
		MenuRouter.POST("getMenuList", sys.GetMenuList)         // 分页获取基础menu列表
		MenuRouter.POST("addBaseMenu", sys.AddBaseMenu)         // 新增菜单
		MenuRouter.POST("getBaseMenuTree", sys.GetBaseMenuTree) // 获取用户动态路由
		//MenuRouter.POST("addMenuAuthority", v1.AddMenuAuthority) //	增加menu和角色关联关系
		MenuRouter.POST("getMenuAuthority", sys.GetMenuAuthority) // 获取指定角色menu
		MenuRouter.POST("deleteBaseMenu", sys.DeleteBaseMenu)     // 删除菜单
		MenuRouter.POST("updateBaseMenu", sys.UpdateBaseMenu)     // 更新菜单
		MenuRouter.POST("getBaseMenuById", sys.GetBaseMenuById)   // 根据id获取菜单
	}
	return MenuRouter
}
