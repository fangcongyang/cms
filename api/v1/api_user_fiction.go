package v1

import (
	"cms/global"
	"cms/model"
	"cms/model/request"
	"cms/model/response"
	"cms/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// @Tags AddBookShelf
// @Summary 添加用户收藏小说信息
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /api/userFiction/save [post]

func AddBookShelf(c *gin.Context) {
	var userFiction model.UserFiction
	_ = c.ShouldBind(&userFiction)
	var userInfo model.UserInfo
	userJson := global.GVA_REDIS.Get(userFiction.TokenId).Val()
	err := json.Unmarshal([]byte(userJson), &userInfo)
	if err != nil {
		return
	}
	userFiction.UserId = userInfo.Id
	err = service.AddBookShelf(userFiction)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.Ok(c)
}

// @Tags ListBookShelf
// @Summary 分页获取用户收藏小说信息
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /api/userFiction/listBookShelf [post]

func ListBookShelf(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBind(&pageInfo)
	var userInfo model.UserInfo
	userJson := global.GVA_REDIS.Get(pageInfo.TokenId).Val()
	err := json.Unmarshal([]byte(userJson), &userInfo)
	if err != nil {
		return
	}
	fictionList, total, err := service.ListBookShelf(userInfo.Id, pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithPage(fictionList, total, pageInfo, c)
}
