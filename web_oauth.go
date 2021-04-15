package hwxsdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ScopeType string

const (
	SNSAPI_BASE     ScopeType = "snsapi_base"
	SNSAPI_USERINFO ScopeType = "snsapi_userinfo"
)

// 获取Code的连接, state值在回调跳转时原样带回
func Oauth2Url(cfg *Config, scope ScopeType, state string, redirectUri string) string {
	return fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect", cfg.Appid, url.QueryEscape(redirectUri), scope, state)
}

type WebAccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

func (o *WebAccessToken) ToJson() string {
	bytes, _ := json.Marshal(&o)
	return string(bytes)
}

type Oauth2Error struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// AccessToken 通过code换取网页授权access_token
// 微信网页授权是通过OAuth2.0机制实现的，在用户授权给公众号后，公众号可以获取到一个网页授权特有的接口调用凭证（网页授权access_token），通过网页授权access_token可以进行授权后接口调用，如获取用户基本信息；
func Oauth2AccessToken(cfg *Config, code string) (t *WebAccessToken, err error) {
	url := fmt.Sprintf("%s/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", cfg.GetServiceUrl(), cfg.Appid, cfg.Secret, code)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &t)
	if err != nil {
		return
	}
	if t != nil {
		return
	}

	var e *Oauth2Error
	err = json.Unmarshal(bytes, &e)
	if err != nil {
		return
	}
	if e != nil {
		err = errors.New(string(e.Errcode) + ":" + e.Errmsg)
		return
	}
	if e == nil {
		err = errors.New("微信响应报文解析异常")
		return
	}
	return
}

// 刷新token
func Oauth2RefreshToken(cfg *Config, refreshToken string) (t *WebAccessToken, err error) {
	url := fmt.Sprintf("%s/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s", cfg.GetServiceUrl(), cfg.Appid, refreshToken)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &t)
	if err != nil {
		return
	}
	if t != nil {
		return
	}

	var e *Oauth2Error
	err = json.Unmarshal(bytes, &e)
	if err != nil {
		return
	}
	if e != nil {
		err = errors.New(string(e.Errcode) + ":" + e.Errmsg)
		return
	}
	if e == nil {
		err = errors.New("微信响应报文解析异常")
		return
	}
	return
}

type WebUserInfo struct {
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int64    `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

func (o *WebUserInfo) ToJson() string {
	b, _ := json.Marshal(&o)
	return string(b)
}

// UserInfo 获取微信用户基础信息
func Oauth2UserInfo(cfg *Config, accessToken string, openid string) (t *WebUserInfo, err error) {
	url := fmt.Sprintf("%s/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", cfg.GetServiceUrl(), accessToken, openid)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &t)
	if err != nil {
		return
	}
	if t != nil {
		return
	}

	var e *Oauth2Error
	err = json.Unmarshal(bytes, &e)
	if err != nil {
		return
	}
	if e != nil {
		err = errors.New(string(e.Errcode) + ":" + e.Errmsg)
		return
	}
	if e == nil {
		err = errors.New("微信响应报文解析异常")
		return
	}
	return
}

// Check 检验授权凭证（access_token）是否有效
func Oauth2Check(cfg *Config, accessToken string, openid string) (ok bool, err error) {
	url := fmt.Sprintf("%s/sns/auth?access_token=%s&openid=%s", cfg.GetServiceUrl(), accessToken, openid)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var e *Oauth2Error
	err = json.Unmarshal(bytes, &e)
	if err != nil {
		return
	}
	if e == nil {
		err = errors.New("微信响应报文解析异常")
		return
	}
	if e.Errcode != 0 {
		err = errors.New(string(e.Errcode) + ":" + e.Errmsg)
		return
	}
	return true, nil
}
