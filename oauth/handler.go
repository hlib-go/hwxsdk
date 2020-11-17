package oauth

import "net/http"

// CallbackHandler 微信授权回调请求handle
// 业务系统根据code处理业务逻辑
func CbFuncHandler(cbFunc func(code, state string, writer http.ResponseWriter, request *http.Request)) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		code := request.FormValue("code")
		state := request.FormValue("state")
		cbFunc(code, state, writer, request)
	})
}
