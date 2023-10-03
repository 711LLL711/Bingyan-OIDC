package oauth2

import (
	"time"
)

// 默认过期时间
const defaultExpiryDelta = 10 * time.Second

type Clientcode struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Code         string
}

type Token struct {
	AccessToken  string
	RefreshToken string
	Expiry       time.Time
	expiryDelta  time.Duration
}

func (t *Token) IsExpired() bool {
	return t != nil && t.AccessToken != "" && t.Expiry.Before(time.Now())
}

// 先只做授权码
type Endpoint struct {
	AuthURL  string
	TokenURL string
}

type Config struct {
	ClientID string //客户端ID

	ClientSecret string //客户端密钥

	Endpoint Endpoint //认证服务器地址

	RedirectURL string //重定向地址

	Scopes []string //授权范围
}

type ServerInfo struct {
	AuthURL  string //认证服务器地址
	AuthType int    //认证类型
}
