package model

import "time"

//用于申请OAuth服务的请求
type ClientInfo struct {
	ClientName   string    //申请服务的客户端名称
	ClientID     string    `gorm:"primaryKey;column:client_id"` //客户端ID
	ClientSecret string    `gorm:"column:client_secret"`        //客户端密钥
	ClientDomain string    `gorm:"column:client_domain"`        //申请OAuth服务的域名
	RedirectURI  string    `gorm:"column:redirect_url"`         //重定向URI
	CreateAt     time.Time `gorm:"column:create_at;type:TIMESTAMP"`
}

func NewDefaultClientInfo() *ClientInfo {
	return &ClientInfo{
		CreateAt: time.Now(),
	}
}

// GetID client id
func (c *ClientInfo) GetID() string {
	return c.ClientID
}

// GetSecret client secret
func (c *ClientInfo) GetSecret() string {
	return c.ClientSecret
}

// GetDomain client domain
func (c *ClientInfo) GetDomain() string {
	return c.ClientSecret
}

// GetRedirectURI redirect URI
func (c *ClientInfo) GetRedirectURI() string {
	return c.RedirectURI
}
