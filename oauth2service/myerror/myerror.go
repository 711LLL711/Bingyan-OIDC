package myerror

import (
	"errors"
)

var (
	ErrInvalidRedirectURI   = errors.New("invalid redirect uri")
	ErrInvalidAuthorizeCode = errors.New("invalid authorize code")
	ErrInvalidAccessToken   = errors.New("invalid access token")
	ErrInvalidRefreshToken  = errors.New("invalid refresh token")
	ErrExpiredAccessToken   = errors.New("expired access token")
	ErrExpiredRefreshToken  = errors.New("expired refresh token")
	ErrUnauthorizedClient   = errors.New("unauthorized_client")
	ErrAccessDenied         = errors.New("access_denied")
	ErrServerError          = errors.New("server_error")
	ErrInvalidClientInfo    = errors.New("invalid CLIENT info")
)

//设置error对应的描述

var Descriptions = map[error]string{
	ErrInvalidRedirectURI:   "invalid redirect uri",
	ErrInvalidAuthorizeCode: "invalid authorize code",
	ErrInvalidAccessToken:   "invalid access token",
	ErrInvalidRefreshToken:  "invalid refresh token",
	ErrExpiredAccessToken:   "expired access token",
	ErrExpiredRefreshToken:  "expired refresh token",
	ErrUnauthorizedClient:   "unauthorized_client",
	ErrAccessDenied:         "The resource owner or authorization server denied the request",
	ErrServerError:          "server_error",
	ErrInvalidClientInfo:    "Either no client_id or client_secrect sent",
}

// 定义不同错误的状态码
var StatusCodes = map[error]int{
	ErrInvalidRedirectURI:   400,
	ErrInvalidAuthorizeCode: 400,
	ErrInvalidAccessToken:   401,
	ErrInvalidRefreshToken:  401,
	ErrExpiredAccessToken:   401,
	ErrExpiredRefreshToken:  401,
	ErrUnauthorizedClient:   401,
	ErrAccessDenied:         403,
	ErrServerError:          500,
}
