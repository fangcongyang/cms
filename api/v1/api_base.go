package v1

import (
	"cms/global"
	"cms/model/response"
	"cms/service"
	"cms/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]

func ApiLogin(c *gin.Context) {
	userInfo, err := service.LoginByPassword(c.PostForm("mobile"), c.PostForm("password"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		tokenId := utils.MD5V([]byte(utils.Int64ToStr(userInfo.Id) + userInfo.Mobile + utils.Int64ToStr(time.Time{}.UnixMilli())))
		userInfo.TokenId = tokenId
		data, _ := json.Marshal(userInfo)
		statusCmd := global.GVA_REDIS.Set(tokenId, data, time.Minute*30)
		fmt.Printf("", statusCmd)
		response.Result(0, userInfo, "登录成功!", c)
	}
}

// @Tags Base
// @Summary 获取banner信息
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]

func ApiBanner(c *gin.Context) {
	bannerType, err := utils.StrToInt(c.PostForm("type"))
	if err != nil {
		bannerType = 1
	}
	bannerList, err := service.SelectBannerList(bannerType)
	response.OkWithData(bannerList, c)
}
