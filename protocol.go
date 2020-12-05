package hwxsdk

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var log = logrus.WithField("sdk", "hwxsdk")

type Error struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (e *Error) Error() string {
	return e.Errmsg + ".[" + strconv.FormatInt(e.Errcode, 10) + "]"
}

func (e *Error) Resp(resp *http.Response) (rb []byte, err error) {
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status + ".[WXERR]")
		return
	}
	rb, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(rb, &e)
	if err != nil {
		return
	}
	if e.Errcode > 0 {
		err = ErrMsg(e)
		return
	}
	return
}

// 微信接口GET请求
func WxGet(c *Config, path string) (rb []byte, err error) {
	var (
		wlog    = log.WithField("requestId", Rand32())
		begTime = time.Now().UnixNano()
		url     = c.GetServiceUrl() + path
	)
	defer func() {
		wlog.Info("微信请求URL: ", url)
		wlog.Info("微信响应Body: ", string(rb), " ", strconv.FormatInt((time.Now().UnixNano()-begTime)/1e6, 10), "ms")
		if err != nil {
			wlog.Warn("微信响应错误： ", err.Error())
		}
	}()
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	rb, err = new(Error).Resp(resp)
	if err != nil {
		return
	}
	return
}

func WxGetUnmarshal(c *Config, path string, result interface{}) (err error) {
	rb, err := WxGet(c, path)
	if err != nil {
		return
	}
	err = json.Unmarshal(rb, result)
	if err != nil {
		return
	}
	return
}

// 微信接口POST请求
func WxPost(c *Config, path string, bm interface{}) (rb []byte, err error) {
	if bm == nil {
		bm = make(map[string]string)
	}
	reqBytes, err := json.Marshal(bm)
	if err != nil {
		return
	}
	var (
		wlog    = log.WithField("requestId", Rand32())
		begTime = time.Now().UnixNano()
		url     = c.GetServiceUrl() + path
		reqBody = string(reqBytes)
	)
	defer func() {
		wlog.Info("微信请求URL: ", url)
		wlog.Info("微信请求报文: ", reqBody)
		wlog.Info("微信响应报文: ", string(rb), " ", strconv.FormatInt((time.Now().UnixNano()-begTime)/1e6, 10), "ms")
		if err != nil {
			wlog.Warn("微信响应错误:", err.Error())
		}
	}()
	resp, err := http.Post(url, "application/json; charset=UTF-8", strings.NewReader(reqBody))
	if err != nil {
		return
	}
	rb, err = new(Error).Resp(resp)
	if err != nil {
		return
	}
	return
}

func WxPostUnmarshal(c *Config, path string, bm interface{}, result interface{}) (err error) {
	rb, err := WxPost(c, path, bm)
	if err != nil {
		return
	}
	err = json.Unmarshal(rb, result)
	if err != nil {
		return
	}
	return
}
