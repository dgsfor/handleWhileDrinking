package middleware

import (
	"github.com/gin-gonic/gin"
	"handleWhileDrinking/serializer"
	"handleWhileDrinking/service/oauth"
	"handleWhileDrinking/util"
	"net/http"
)

// 登录认证拦截器
func MiddlewareAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := util.GetUserCookie(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, serializer.Response{
				Code:  http.StatusUnauthorized,
				Data:  nil,
				Msg:   "认证失败",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

// 微信登录认证拦截器
func MiddlewareWxAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		rdSession := c.Request.Header.Get("3rd_session")
		if !oauth.JuggleInRedis(rdSession) {
			c.JSON(http.StatusUnauthorized, serializer.SreResponse{
				Response: serializer.Response{
					Code: http.StatusUnauthorized,
					Data: nil,
					Msg:  "token失效",
				},
				ReCode: serializer.PARAMS_ERROR,
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}