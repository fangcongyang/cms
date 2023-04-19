package v1

import (
	"cms/global"
	"cms/model/request"
	"cms/model/response"
	"cms/service"
	"cms/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]

func QueryFictionSortPageList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBind(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err, total := service.QueryFictionSortPageList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithPage(list, total, pageInfo, c)
	}
}

//@author: fcy
//@function: Get
//@description: 通过id获取小说信息
//@param: id int64
//@return: err error, user *model.SysUser

func GetFictionById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	m, err := service.GetFictionById(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(m, c)
}

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]

func QueryFictionPageList(c *gin.Context) {
	var fictionApiParam request.FictionApiParam
	_ = c.ShouldBind(&fictionApiParam)
	if err := utils.Verify(fictionApiParam, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err, total := service.QueryFictionPageList(fictionApiParam); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			Records: list,
			Total:   total,
			Current: fictionApiParam.PageNo,
			Size:    fictionApiParam.PageSize,
			Pages:   1,
		}, "获取成功", c)
	}
}

//@author: fcy
//@function: Get
//@description: 通过id获取小说信息
//@param: id int64
//@return: err error, user *model.SysUser

func GetFictionChapterById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	m, err := service.GetFictionChapterById(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(m, c)
}

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]

func QueryFictionChapterList(c *gin.Context) {
	fictionId, _ := strconv.ParseInt(c.Query("fictionId"), 10, 64)
	if list, err := service.QueryFictionChapterPageList(fictionId); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}
