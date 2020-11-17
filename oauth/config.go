package oauth

import "encoding/json"

// 微信网页授权配置参数
type Config struct {
	Oauth2Url   string `json:"oauth2Url"`
	ServiceUrl  string `json:"serviceUrl"`
	Appid       string `json:"appid"`
	Secret      string `json:"secret"`
	RedirectUri string `json:"redirectUri"` //不带参数的URL链接，用于接收微信code
}

func NewConfig(appid, secret, redirectUri string) *Config {
	return &Config{
		Oauth2Url:   "https://open.weixin.qq.com",
		ServiceUrl:  "https://api.weixin.qq.com",
		Appid:       appid,
		Secret:      secret,
		RedirectUri: redirectUri,
	}
}

func (c *Config) JSON() string {
	b, _ := json.Marshal(c)
	return string(b)
}
