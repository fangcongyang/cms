package middleware

import (
	"cms/consts"
	"cms/global"
	"cms/model/response"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		requestUrl := c.Request.RequestURI
		if strings.HasPrefix(requestUrl, "/api") && !strings.Contains(requestUrl, "/api/login") &&
			!strings.HasPrefix(requestUrl, "/filed") && !strings.HasPrefix(requestUrl, "/static") && !strings.HasPrefix(requestUrl, "/api/es") {
			var tokenId string
			if c.Request.Method == "GET" {
				tokenId = c.Query("tokenId")
			} else {
				tokenId = c.PostForm("tokenId")
			}
			if tokenId == "" {
				response.Result(consts.NOT_LOGIN, gin.H{"reload": true}, "未登录或非法访问", c)
				c.Abort()
				return
			}
			userJson := global.GVA_REDIS.Get(tokenId).Val()
			global.GVA_REDIS.Set(tokenId, userJson, 30*time.Minute)
		}
		c.Next()
	}
}
