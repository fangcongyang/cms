package router

import (
	"cms/router/sys"
	"github.com/gin-gonic/gin"
)

func InitSysPublicRouter(router *gin.RouterGroup) {
	sys.InitBaseRouter(router) // 注册基础功能路由 不做鉴权
	sys.InitInitRouter(router) // 自动初始化相关
}

func InitSysPrivateRouter(router *gin.RouterGroup) {
	sys.InitApiRouter(router)                 // 注册功能api路由
	sys.InitJwtRouter(router)                 // jwt相关路由
	sys.InitUserRouter(router)                // 注册用户路由
	sys.InitMenuRouter(router)                // 注册menu路由
	sys.InitEmailRouter(router)               // 邮件相关路由
	sys.InitSystemRouter(router)              // system相关路由
	sys.InitCasbinRouter(router)              // 权限相关路由
	InitCustomerRouter(router)                // 客户路由
	sys.InitAutoCodeRouter(router)            // 创建自动化代码
	sys.InitAuthorityRouter(router)           // 注册角色路由
	InitSimpleUploaderRouter(router)          // 断点续传（插件版）
	sys.InitSysDictionaryRouter(router)       // 字典管理
	sys.InitSysOperationRecordRouter(router)  // 操作记录
	sys.InitSysDictionaryDetailRouter(router) // 字典详情管理
	InitFileUploadAndDownloadRouter(router)   // 文件上传下载功能路由
	InitExcelRouter(router)                   // 表格导入导出
}
