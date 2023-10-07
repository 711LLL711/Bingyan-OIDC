package response

import (
	"net/http"
)

type ResponseType string

// 回复的是授权码还是令牌
const (
	AuthorizationCode ResponseType = "code"
	AccessToken       ResponseType = "token" //令牌
)

func (t ResponseType) GetAuthorizeCode() string {
	return string(t)
}

type Response struct {
	Header     http.Header
	Error      error
	Hint       string
	StatusCode int
	Data       interface{}
}

type ResponseBody struct {
	Success bool
	Hint    string
	Data    interface{}
}

// 创建请求
func NewResponse(err error, statusCode int) *Response {
	return &Response{
		Error:      err,
		StatusCode: statusCode,
	}
}

// 设置请求头
func (r *Response) SetHeader(key, value string) {
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	r.Header.Set(key, value)
}

func MakeResponseBody(state bool, hint string, data interface{}) ResponseBody {
	return ResponseBody{
		Success: state,
		Hint:    hint,
		Data:    data,
	}
}

func MakeFailedResponseBody(hint string) ResponseBody {
	return MakeResponseBody(false, hint, nil)
}
func MakeSuccessResponseBody(hint string, data interface{}) ResponseBody {
	return MakeResponseBody(true, hint, data)
}
