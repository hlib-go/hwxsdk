package main

import (
	"fmt"
	"github.com/hlib-go/hwxsdk/oauth"
	"log"
	"net/http"
)

// 微信授权测试
// https://msd.himkt.cn/work/hwwxsdk/oauth/base
// https://msd.himkt.cn/work/hwwxsdk/oauth/userinfo
func main() {
	appid := "wx239c521c61221a8a"
	secret := "******"

	redirectUriPath := "/hwxsdk/oauth/cburl"
	var cfg = oauth.NewConfig(appid, secret, "https://msd.himkt.cn/work"+redirectUriPath)
	log.Println("cfg=", cfg.JSON())

	// 跳转微信授权 SNSAPI_BASE
	http.HandleFunc("/hwwxsdk/oauth/base", func(writer http.ResponseWriter, request *http.Request) {
		url := oauth.Url(cfg, oauth.SNSAPI_BASE, "snsapi_base")
		writer.Header().Set("Location", url)
		writer.WriteHeader(302)
	})
	// 跳转微信授权 SNSAPI_USERINFO
	http.HandleFunc("/hwwxsdk/oauth/userinfo", func(writer http.ResponseWriter, request *http.Request) {
		url := oauth.Url(cfg, oauth.SNSAPI_USERINFO, "snsapi_userinfo")
		writer.Header().Set("Location", url)
		writer.WriteHeader(302)
	})

	// 授权后的回调页面
	// https://msd.himkt.cn/work/hwxsdk/oauth/cburl
	http.Handle(redirectUriPath, oauth.CbFuncHandler(func(code, state string, writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html;charset=UTF-8")
		var (
			err error
			at  *oauth.Oauth2AccessToken
			u   *oauth.Oauth2UserInfo
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

		at, err = oauth.AccessToken(cfg, code)
		if err != nil {
			log.Println("oauth.AccessToken error:", err.Error())
			return
		}
		log.Println("AccessToken", at.ToJson())
		rt, err := oauth.RefreshToken(cfg, at.RefreshToken)
		if err != nil {
			log.Println("oauth.RefreshToken error:", err.Error())
			return
		}
		log.Println("RefreshToken", rt.ToJson())
		if state == "snsapi_userinfo" {
			u, err = oauth.UserInfo(cfg, at.AccessToken, at.Openid)
			if err != nil {
				log.Println("oauth.UserInfo error:", err.Error())
				return
			}
			log.Println("UserInfo", u.ToJson())
		}
		ok, err := oauth.Check(cfg, at.AccessToken, at.Openid)
		if err != nil {
			log.Println("oauth.Check error:", err.Error())
			return
		}
		log.Println("oauth.Check ", ok)
	}))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 80), nil))
}
