package util

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

// Cookie Names
const (
	UserCookieName      = "ssopa_passport"
	SignatureCookieName = "ssopa_signature"
	CookieExpiration    = 60 * 60 * 24 * 7 // 1 week
)

// UserCookie the truely info in cookie
type UserCookie struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Expiration int64  `json:"expiration"`
}


// IsExpired return true when cookie expired
func (u *UserCookie) IsExpired() bool {
	return u.Expiration < time.Now().Unix()
}

// SetLoginCookies in response
func SetLoginCookies(c *gin.Context, username string,email string) (err error) {
	// set user cookie
	userCookie := &UserCookie{
		Name:       username,
		Email:      email,
		Expiration: time.Now().Unix() + CookieExpiration,
	}
	encodedValue, err := json.Marshal(userCookie)
	if err != nil {
		return err
	}
	c.SetCookie(UserCookieName, base64.StdEncoding.EncodeToString(encodedValue), CookieExpiration, "/", "", false, false)
	// set signature
	sig, err := SignData(encodedValue)
	if err != nil {
		return err
	}
	c.SetCookie(SignatureCookieName, base64.StdEncoding.EncodeToString(sig), CookieExpiration, "/", "", false, false)
	return err
}

// Logout reset client cookie in response, set max age to 0 actually
// due to Set-Cookie will contain other cookies so DO NOT use header.Del("Set-Cookie")
func Logout(c *gin.Context) {
	c.SetCookie(UserCookieName, "", -1, "/", "", false, false)
	c.SetCookie(SignatureCookieName, "", -1, "/", "", false, false)
}

// GetUserCookie from request header
func GetUserCookie(c *gin.Context) (u *UserCookie, err error) {
	userInfo, err := c.Cookie(UserCookieName)
	if err != nil {
		return nil, err
	}
	// base64 decode
	decodedUserInfo, err := base64.StdEncoding.DecodeString(userInfo)
	if err != nil {
		return nil, err
	}

	// get signature
	sig, err := c.Cookie(SignatureCookieName)
	if err != nil {
		return nil, err
	}
	decodedSignature, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		return nil, err
	}

	err = VerifyData(decodedUserInfo, decodedSignature)
	if err != nil {
		// signature verify failed
		return nil, err
	}

	// decode cookie
	u = &UserCookie{}
	if err = json.Unmarshal(decodedUserInfo, u); err != nil {
		return
	}
	if u.IsExpired() {
		return nil, errors.New("cookie expired")
	}
	return u, nil
}
