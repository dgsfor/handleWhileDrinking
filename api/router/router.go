package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	oauthController "handleWhileDrinking/controller/oauth"
	"handleWhileDrinking/middleware"
	"io/ioutil"
	"time"
)

var (
	Logger, _ = zap.NewProduction()
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	atLeastLoginMiddleware := middleware.MiddlewareAuthorization()
	atLeastWxLoginMiddleware := middleware.MiddlewareWxAuthorization()
	r := gin.Default()
	r.Use(middleware.Cors())
	// 关闭gin默认AccessLog日志
	if gin.Mode() == "release" {
		gin.DefaultWriter = ioutil.Discard
		r.Use(middleware.GinZap(Logger, time.RFC3339, true))
		r.Use(gin.Recovery())
	}
	// 路由
	v1 := r.Group("/api/hwd/v1")
	{
		// user
		ur := v1.Group("/oauth")
		ur.POST("/qywx/login", oauthController.QywxLogin)
		ur.GET("/qywx/logout", oauthController.QywxLogout)
		ur.GET("/qywx/check_login", atLeastLoginMiddleware, oauthController.CheckLogin)
		ur.POST("/wx/login",oauthController.WxLogin)
		ur.GET("/wx/logout",oauthController.WxLogout)
		ur.GET("/wx/check_login", atLeastWxLoginMiddleware, oauthController.CheckLogin)
	}
	return r
}
