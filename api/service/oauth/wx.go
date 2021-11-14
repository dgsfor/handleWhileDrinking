package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"handleWhileDrinking/conf"
	"handleWhileDrinking/serializer"
	"handleWhileDrinking/util"
	"io/ioutil"
	"net/http"
)

type WxLoginParams struct {
	// 用户登录凭证（有效期五分钟）。开发者需要在开发者服务器后台调用 api，使用 code 换取 userid 和 session_key 等信息
	WxLoginCode string `form:"wx_login_code" json:"wx_login_code" binding:"required"`
}

func (p *WxLoginParams) WxLogin() serializer.SreResponse {
	// 重新获取session
	result, err := GetWxCode2Session(p.WxLoginCode)
	if err != nil {
		return serializer.SreResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取微信code2Session失败",
			},
			ReCode: serializer.COM_ERROR,
		}
	}
	rdSession := Get3rdSession()
	// 写入session到redis防止频繁登录
	wxSu := SetInRedis(rdSession, result.Openid+";"+result.SessionKey, 2999)
	if !wxSu {
		return serializer.SreResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: "set redis error",
				Msg:  "redis操作失败",
			},
			ReCode: serializer.COM_ERROR,
		}
	}
	return serializer.SreResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: rdSession,
			Msg:  "登录成功",
		},
		ReCode: serializer.ALL_SUCCESS,
	}
}

func WxLogout(rdSession string) serializer.SreResponse  {
	wxSu := DelInRedis(rdSession)
	if !wxSu {
		return serializer.SreResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: "del in redis error",
				Msg:  "退出登录失败",
			},
			ReCode: serializer.COM_ERROR,
		}
	}
	return serializer.SreResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "退出登录成功",
		},
		ReCode: serializer.ALL_SUCCESS,
	}
}

func Get3rdSession() string {
	sessionId, _ := uuid.NewRandom()
	return util.HashAndSalt([]byte(sessionId.String()))
}

/**
https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
*/
type WxCode2Session struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int64  `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func GetWxCode2Session(wxLoginCode string) (result *WxCode2Session, err error) {
	retry := 0
Retry:
	resp, err := httpClient.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_codes",
		conf.GetConfig("wx::app_id").String(),
		conf.GetConfig("wx::app_secret").String(),
		wxLoginCode))
	if err != nil {
		if retry < retries {
			retry++
			goto Retry
		}
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result = &WxCode2Session{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
