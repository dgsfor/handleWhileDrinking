package oauth

import (
	"github.com/gin-gonic/gin"
	"handleWhileDrinking/serializer"
	oauth_service "handleWhileDrinking/service/oauth"
	"net/http"
)


// qywx login
func QywxLogin(c *gin.Context) {
	svc := oauth_service.QywxLoginParams{}
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
	result := svc.QywxLogin()
	c.JSON(result.Code, serializer.SreResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ReCode: result.ReCode,
	})
}

// qywx logout
func QywxLogout(c *gin.Context) {
	svc := oauth_service.QywxLogoutParams{}
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
	result := svc.QywxLogout()
	c.JSON(result.Code, serializer.SreResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ReCode: result.ReCode,
	})
}


func CheckLogin(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}