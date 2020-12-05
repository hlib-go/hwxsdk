package test

import (
	"github.com/hlib-go/hwxsdk"
	"testing"
)

func GetAccessToken(cfg *hwxsdk.Config) string {
	token, err := hwxsdk.WxAccessToken(cfg)
	if err != nil {
		return ""
	}
	return token.AccessToken
}

func TestAccessToken(t *testing.T) {
	token, err := hwxsdk.WxAccessToken(cfg)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("token.AccessTokent", token.AccessToken)
}
