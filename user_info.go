package hwxsdk

import (
	"fmt"
)

// 用户管理：获取用户基本信息 https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
func WxUserInfo(c *Config, openid string) (user *UserInfo, err error) {
	// GET /cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
	path := fmt.Sprintf("/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN", c.AccessToken, openid)
	err = WxGetUnmarshal(c, path, &user)
	if err != nil {
		return
	}
	return
}

type UserInfo struct {
	Subscribe      int64   `json:"subscribe"` // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	Openid         string  `json:"openid"`
	Nickname       string  `json:"nickname"`
	Sex            int64   `json:"sex"`
	City           string  `json:"city"`
	Country        string  `json:"country"`
	Province       string  `json:"province"`
	Language       string  `json:"language"`
	Headimgurl     string  `json:"headimgurl"`
	SubscribeTime  int64   `json:"subscribe_time"`
	Unionid        string  `json:"unionid"`
	Remark         string  `json:"remark"`
	Groupid        int64   `json:"groupid"`
	TagidList      []int64 `json:"tagid_list"`
	SubscribeScene string  `json:"subscribe_scene"`
	QrScene        int64   `json:"qr_scene"`
	QrScene_str    string  `json:"qr_scene_str"`
}
