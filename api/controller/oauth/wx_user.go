package oauth

import (
	"github.com/gin-gonic/gin"
	"handleWhileDrinking/serializer"
	oauth_service "handleWhileDrinking/service/oauth"
	"net/http"
)

// wx login
func WxLogin(c *gin.Context) {
	svc := oauth_service.WxLoginParams{}
	//rdSession := c.Request.Header.Get("3rd_session")
	if err := c.ShouldBind(&svc); err != nil {
		c.JSON(http.StatusBadRequest, serializer.SreResponse{
			Response: serializer.Response{
				Code: http.StatusBadRequest,
				Data: err.Error(),
				Msg:  "参数错误，请检查！",
			},
			ReCode: serializer.PARAMS_ERROR,
		})
		return
	}
	result := svc.WxLogin()
	c.JSON(result.Code, serializer.SreResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ReCode: result.ReCode,
	})
}

// wx logout
func WxLogout(c *gin.Context) {
	rdSession := c.Request.Header.Get("3rd_session")
	result := oauth_service.WxLogout(rdSession)
	c.JSON(result.Code, serializer.SreResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ReCode: result.ReCode,
	})
}
