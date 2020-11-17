package main

import (
	"github.com/hlib-go/hwxsdk/oauth"
	"testing"
)

var cfg = oauth.NewConfig("wx239c521c61221a8a", "86770be0bde9017130e195e87a471509", "https://msd.himkt.cn/work/hwxsdk/oauth/cburl")

func TestUrl(t *testing.T) {
	url := oauth.Url(cfg, oauth.SNSAPI_BASE, "")
	t.Log(url)
}
