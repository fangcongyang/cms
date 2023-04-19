package response

import (
	"cms/consts"
	"cms/model/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PageResult struct {
	Records interface{} `json:"records"`
	Total   int64       `json:"total"`
	Current int         `json:"current"`
	Size    int         `json:"size"`
	Pages   int64       `json:"pages"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(consts.SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(consts.SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(consts.SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(consts.SUCCESS, data, message, c)
}

func OkWithPage(data interface{}, total int64, pageInfo request.PageInfo, c *gin.Context) {
	var pages int64
	if total%int64(pageInfo.PageSize) == 0 {
		pages = total / int64(pageInfo.PageSize)
	} else {
		pages = total/int64(pageInfo.PageSize) + 1
	}
	page := PageResult{
		Records: data,
		Total:   total,
		Current: pageInfo.PageNo,
		Size:    pageInfo.PageSize,
		Pages:   pages,
	}
	OkWithData(page, c)
}

func Fail(c *gin.Context) {
	Result(consts.FAIL, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(consts.FAIL, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(consts.FAIL, data, message, c)
}
