package model

import (
	"time"
)

// 新建token变量并返回指针
func NewToken() *Token {
	return &Token{}
}

type Token struct {
	ClientID         string    `gorm:"foreignKey;column:client_id"` //链接到clientinfo的clientid
	UserID           string    `gorm:"column:user_id"`
	RedirectURI      string    `gorm:"column:redirect_url"`
	Scope            string    `gorm:"column:scope"`
	Code             string    `gorm:"column:code"` //授权码
	CodeCreateAt     time.Time `gorm:"column:code_create_at;type:DATETIME"`
	CodeExpiresIn    time.Time `gorm:"column:xode_expire_in"` //time.duration要转换成string存储
	Access           string    `gorm:"column:access"`         //令牌
	AccessCreateAt   time.Time `gorm:"column:access_create_at;type:DATETIME"`
	AccessExpiresIn  time.Time `gorm:"column:access_expire_in"`
	Refresh          string    `gorm:"column:refresh"`
	RefreshCreateAt  time.Time `gorm:"column:refresh_create_at;type:DATETIME"`
	RefreshExpiresIn time.Time `gorm:"column:refresh_expire_in"`
}

type QueryTokenRequest struct {
	ClientID     string
	ClientSecret string
	Code         string
}

// 默认过期时间
var (
	DefaultScope                 = "read"
	DefaultCodeCreateAt          = time.Now()
	DefaultAccessCreateAt        = time.Now()
	DefaultCodeExpireDuration    = time.Minute * 10
	DefaultAccessExpireDuration  = time.Hour * 2
	DefaultRefreshCreateAt       = time.Now()
	DefaultRefreshExpireDuration = time.Hour * 24 * 3
)

// 设置默认的token config
func NewDefaultToken() *Token {
	return &Token{
		Scope:            DefaultScope,
		CodeCreateAt:     DefaultCodeCreateAt,
		AccessCreateAt:   DefaultAccessCreateAt,
		CodeExpiresIn:    DefaultCodeCreateAt.Add(DefaultCodeExpireDuration),
		AccessExpiresIn:  DefaultAccessCreateAt.Add(DefaultAccessExpireDuration),
		RefreshExpiresIn: DefaultAccessCreateAt.Add(DefaultRefreshExpireDuration),
		RefreshCreateAt:  DefaultRefreshCreateAt,
	}
}

// refreshtoken
func RefreshToken(token *Token) {
	token.AccessCreateAt = time.Now()
	token.AccessExpiresIn = time.Now().Add(DefaultAccessExpireDuration)
}

//设置或获得token的值

// GetClientID the client id
func (t *Token) GetClientID() string {
	return t.ClientID
}

// SetClientID the client id
func (t *Token) SetClientID(clientID string) {
	t.ClientID = clientID
}

// GetRedirectURI redirect URI
func (t *Token) GetRedirectURI() string {
	return t.RedirectURI
}

// SetRedirectURI redirect URI
func (t *Token) SetRedirectURI(redirectURI string) {
	t.RedirectURI = redirectURI
}

// GetScope get scope of authorization
func (t *Token) GetScope() string {
	return t.Scope
}

// SetScope get scope of authorization
func (t *Token) SetScope(scope string) {
	t.Scope = scope
}

// GetCode authorization code
func (t *Token) GetCode() string {
	return t.Code
}

// SetCode authorization code
func (t *Token) SetCode(code string) {
	t.Code = code
}

// GetCodeCreateAt create Time
func (t *Token) GetCodeCreateAt() time.Time {
	return t.CodeCreateAt
}

// SetCodeCreateAt create Time
func (t *Token) SetCodeCreateAt(createAt time.Time) {
	t.CodeCreateAt = createAt
}

// GetCodeExpiresIn the lifetime in seconds of the authorization code
func (t *Token) GetCodeExpiresIn() time.Duration {
	expire, _ := time.ParseDuration(t.CodeExpiresIn.String())
	return expire
}

// SetCodeExpiresIn the lifetime in seconds of the authorization code
func (t *Token) SetCodeExpiresIn(exp time.Time) {
	t.CodeExpiresIn = exp
}

// GetAccess access Token
func (t *Token) GetAccess() string {
	return t.Access
}

// SetAccess access Token
func (t *Token) SetAccess(access string) {
	t.Access = access
}

// GetAccessCreateAt create Time
func (t *Token) GetAccessCreateAt() time.Time {
	return t.AccessCreateAt
}

// SetAccessCreateAt create Time
func (t *Token) SetAccessCreateAt(createAt time.Time) {
	t.AccessCreateAt = createAt
}

// GetAccessExpiresIn the lifetime in seconds of the access token
func (t *Token) GetAccessExpiresIn() time.Time {
	return t.AccessExpiresIn
}

// SetAccessExpiresIn the lifetime in seconds of the access token
func (t *Token) SetAccessExpiresIn(exp time.Time) {
	t.AccessExpiresIn = exp
}

// GetRefresh refresh Token
func (t *Token) GetRefresh() string {
	return t.Refresh
}

// SetRefresh refresh Token
func (t *Token) SetRefresh(refresh string) {
	t.Refresh = refresh
}

// GetRefreshCreateAt create Time
func (t *Token) GetRefreshCreateAt() time.Time {
	return t.RefreshCreateAt
}

// SetRefreshCreateAt create Time
func (t *Token) SetRefreshCreateAt(createAt time.Time) {
	t.RefreshCreateAt = createAt
}

// GetRefreshExpiresIn the lifetime in seconds of the refresh token
func (t *Token) GetRefreshExpiresIn() time.Time {
	return t.RefreshExpiresIn
}

// SetRefreshExpiresIn the lifetime in seconds of the refresh token
func (t *Token) SetRefreshExpiresIn(exp time.Time) {
	t.RefreshExpiresIn = exp
}
