package test

import (
	"encoding/json"
	"github.com/hlib-go/hwxsdk"
	"testing"
)

func TestUserInfo(t *testing.T) {
	var accessToken = GetAccessToken(cfg)
	user, err := hwxsdk.WxUserInfo(cfg, "", accessToken)
	if err != nil {
		t.Error(err)
		return
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("UserInfo：", string(bytes))
}
