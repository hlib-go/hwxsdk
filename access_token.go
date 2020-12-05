package hwxsdk

import "fmt"

// https请求方式: GET https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
func WxAccessToken(cfg *Config) (token *AccessToken, err error) {
	path := fmt.Sprintf("/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", cfg.AppId, cfg.AppSecret)
	err = WxGetUnmarshal(cfg, path, &token)
	if err != nil {
		return
	}
	return
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
