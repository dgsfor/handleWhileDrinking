package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"handleWhileDrinking/conf"
	"handleWhileDrinking/serializer"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	httpClient = &http.Client{Timeout: 3 * time.Second}
	retries    = 3
)

type QywxLoginParams struct {
	// 用户登录凭证（有效期五分钟）。开发者需要在开发者服务器后台调用 api，使用 code 换取 userid 和 session_key 等信息
	QyLoginCode string `form:"qy_login_code" json:"qy_login_code" binding:"required"`
}

type QywxLogoutParams struct {
	UserId    string `form:"user_id" json:"user_id" binding:"required"`
	UserToken string `form:"user_token" json:"user_token" binding:"required"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"` // access_token 每个用户各不相同,可用于请求api接口,在expires_in内有效,有效期为7200秒
	ExpiresIn   uint64 `json:"expires_in"`   // access_token 过期时间戳
}

type Code2Session struct {
	CorpId string `json:"corpid"` // 企业id
	Userid string `json:"userid"` // 用户id
}

type OAUserData struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data OAUserDataDetail `json:"data"`
}

type OAUserDataDetail struct {
	UserName string `json:"user_name"` //姓名
	Email    string `json:"email"`     //邮箱
	Member   string `json:"member"`    //工号
	Duty     string `json:"duty"`      //岗位
	Dept     string `json:"dept"`      //部门
	Avatar   string `json:"avatar"`    // 头像
}

type ReturnData struct {
	UserToken string           `json:"user_token"` // 返回给小程序端用户态token
	Data      OAUserDataDetail `json:"data"`       // 返回给小程序端用户态数据
	UserId    string           `json:"user_id"`    // 用户id

}

func (p *QywxLoginParams) QywxLogin() serializer.SreResponse {
	su := JuggleInRedis("qywx_access_token")
	if !su {
		// 获取accessToken
		t, err := GetQywxAccessToken()
		if err != nil {
			return serializer.SreResponse{
				Response: serializer.Response{
					Code: http.StatusInternalServerError,
					Data: err,
					Msg:  "获取企业微信access token失败",
				},
				ReCode: serializer.COM_ERROR,
			}
		}
		// 写入accessToken到redis防止频繁登录
		su = SetInRedis("qywx_access_token", t.AccessToken, 2999)
		if !su {
			return serializer.SreResponse{
				Response: serializer.Response{
					Code: http.StatusInternalServerError,
					Data: "set redis error",
					Msg:  "redis操作失败",
				},
				ReCode: serializer.COM_ERROR,
			}
		}
	}
	accessToken, _ := GetInRedis("qywx_access_token")
	fmt.Println(accessToken)
	// 根据accessToken和code获取userid
	userData, err := GetQywxCode2Session(accessToken, p.QyLoginCode)
	if err != nil {
		return serializer.SreResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err,
				Msg:  "获取企业微信userdata失败",
			},
			ReCode: serializer.COM_ERROR,
		}
	}
	su = SetInRedis(userData.Userid, accessToken, 7100)
	if !su {
		return serializer.SreResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: nil,
				Msg:  "redis操作失败",
			},
			ReCode: serializer.COM_ERROR,
		}
	}

	if userData.CorpId != conf.GetConfig("qywx::corpid").String() {
		_ = DelInRedis(userData.Userid)
		return serializer.SreResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err,
				Msg:  "你不是该企业用户",
			},
			ReCode: serializer.COM_ERROR,
		}
	}

	returnData := new(ReturnData)
	//returnData.Data = OAUserData.Data
	//returnData.UserToken = userToken
	returnData.UserId = userData.Userid
	return serializer.SreResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: returnData,
			Msg:  "登录成功",
		},
		ReCode: serializer.ALL_SUCCESS,
	}
}

func (p *QywxLogoutParams) QywxLogout() serializer.SreResponse {
	_ = DelInRedis(p.UserToken)
	accessToken, _ := GetInRedis(p.UserId)
	_ = DelInRedis(accessToken)
	_ = DelInRedis(p.UserId)
	return serializer.SreResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "退出登录成功",
		},
		ReCode: serializer.ALL_SUCCESS,
	}
}

func GetQywxCode2Session(accessToken string, js_code string) (userData *Code2Session, err error) {
	retry := 0
Retry:
	resp, err := httpClient.Get(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/miniprogram/jscode2session?access_token=%s&js_code=%s&grant_type=authorization_code",
		accessToken,
		js_code))
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
	userData = &Code2Session{}
	if err = json.Unmarshal(body, &userData); err != nil {
		return nil, err
	}
	return userData, nil
}

func GetQywxAccessToken() (token *AccessToken, err error) {
	retry := 0
Retry:
	resp, err := httpClient.Get(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s",
		conf.GetConfig("qywx::corpid").String(),
		conf.GetConfig("qywx::suite_secret").String()))
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
	token = &AccessToken{}
	if err = json.Unmarshal(body, &token); err != nil {
		return nil, err
	}
	return token, nil
}

func SetInRedis(key string, value string, expire int) bool {
	rc := conf.RedisClient.Get()
	defer rc.Close()
	_, err := rc.Do("Set", key, value, "EX", expire)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func DelInRedis(key string) bool {
	rc := conf.RedisClient.Get()
	defer rc.Close()
	_, err := rc.Do("Del", key)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func JuggleInRedis(key string) bool {
	rc := conf.RedisClient.Get()
	defer rc.Close()
	ok, err := redis.Bool(rc.Do("exists", key))
	if err != nil {
		return false
	}
	return ok
}

func GetInRedis(key string) (string, bool) {
	rc := conf.RedisClient.Get()
	defer rc.Close()
	t, err := redis.String(rc.Do("GET", key))
	if err != nil {
		return "", false
	}
	return t, true
}
