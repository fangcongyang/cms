package sys

import (
	"cms/api/v1/sys"
	"cms/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		UserRouter.POST("register", sys.Register)
		UserRouter.POST("changePassword", sys.ChangePassword)     // 修改密码
		UserRouter.POST("getUserList", sys.GetUserList)           // 分页获取用户列表
		UserRouter.POST("setUserAuthority", sys.SetUserAuthority) // 设置用户权限
		UserRouter.DELETE("deleteUser", sys.DeleteUser)           // 删除用户
		UserRouter.PUT("setUserInfo", sys.SetUserInfo)            // 设置用户信息
	}
}
