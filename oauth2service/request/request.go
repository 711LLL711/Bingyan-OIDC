package request

import (
	"OIDC/oauth2service/model"
	"OIDC/oauth2service/response"
	"net/http"
)

// 用于申请授权码/令牌的请求
type RequestType struct {
	Response_type response.ResponseType //申请授权码or令牌
	ClientInfo    model.ClientInfo      //客户端信息
	Scope         string                //申请的权限范围
	State         string                //客户端状态，用于防止CSRF攻击
	Request       *http.Request         //http请求
}
