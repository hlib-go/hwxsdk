package hwxsdk

//公众号每次调用接口时，可能获得正确或错误的返回码，开发者可以根据返回码信息调试接口，排查错误。
//全局返回码说明如下：

// 微信返回错误内容改为中文
func ErrMsg(e *Error) (err error) {
	var msg = ""
	switch e.Errcode {
	case 40003:
		msg = "无效openid"
	default:
		msg = e.Errmsg
	}
	e.Errmsg = msg
	return e
}
