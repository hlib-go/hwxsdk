package test

import (
	"fmt"
	"github.com/hlib-go/hwxsdk"
	"log"
	"net/http"
	"testing"
)

// 微信授权测试
// https://msd.himkt.cn/work/hwwxsdk/oauth/base
// https://msd.himkt.cn/work/hwwxsdk/oauth/userinfo
func TestWebOauth(t *testing.T) {
	redirectUriBase := "https://msd.himkt.cn/work"
	redirectUriPath := "/hwxsdk/oauth/cburl"

	// 跳转微信授权 SNSAPI_BASE
	http.HandleFunc("/hwwxsdk/oauth/base", func(writer http.ResponseWriter, request *http.Request) {
		url := hwxsdk.Oauth2Url(cfg, hwxsdk.SNSAPI_BASE, "snsapi_base", redirectUriBase+redirectUriPath)
		writer.Header().Set("Location", url)
		writer.WriteHeader(302)
	})
	// 跳转微信授权 SNSAPI_USERINFO
	http.HandleFunc("/hwwxsdk/oauth/userinfo", func(writer http.ResponseWriter, request *http.Request) {
		url := hwxsdk.Oauth2Url(cfg, hwxsdk.SNSAPI_USERINFO, "snsapi_userinfo", redirectUriBase+redirectUriPath)
		writer.Header().Set("Location", url)
		writer.WriteHeader(302)
	})

	// 授权后的回调页面
	// https://msd.himkt.cn/work/hwxsdk/oauth/cburl
	http.HandleFunc(redirectUriPath, func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html;charset=UTF-8")
		var (
			err   error
			code  = request.FormValue("code")
			state = request.FormValue("state")
			at    *hwxsdk.WebAccessToken
			u     *hwxsdk.WebUserInfo
		)
		defer func() {
			if err != nil {
				writer.Write([]byte(err.Error()))
				return
			}
			writer.Write([]byte("<br>scope:" + state))
			writer.Write([]byte("<br>AccessToken:" + at.ToJson()))
			writer.Write([]byte("<br>UserInfo:" + u.ToJson()))
		}()

		at, err = hwxsdk.Oauth2AccessToken(cfg, code)
		if err != nil {
			log.Println("oauth.AccessToken error:", err.Error())
			return
		}
		log.Println("AccessToken", at.ToJson())
		rt, err := hwxsdk.Oauth2RefreshToken(cfg, at.RefreshToken)
		if err != nil {
			log.Println("oauth.RefreshToken error:", err.Error())
			return
		}
		log.Println("RefreshToken", rt.ToJson())
		if state == "snsapi_userinfo" {
			u, err = hwxsdk.Oauth2UserInfo(cfg, at.AccessToken, at.Openid)
			if err != nil {
				log.Println("oauth.UserInfo error:", err.Error())
				return
			}
			log.Println("UserInfo", u.ToJson())
		}
		ok, err := hwxsdk.Oauth2Check(cfg, at.AccessToken, at.Openid)
		if err != nil {
			log.Println("oauth.Check error:", err.Error())
			return
		}
		log.Println("oauth.Check ", ok)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 80), nil))
}
