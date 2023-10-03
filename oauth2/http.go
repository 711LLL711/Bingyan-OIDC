package oauth2

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	code         string //授权码
	MyToken      string //令牌
	clientconfig = &Clientcode{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:8080/oauth/redirect",
	}
	AuthServerInfo = &ServerInfo{
		AuthURL:  "https://github.com/login/oauth/access_token",
		AuthType: 1,
	}
)

// 设置路由
func SetupRute() {
	c := gin.Default()
	c.GET("/oauth/redirect", func(c *gin.Context) {
		// 从认证服务器获取授权码
		code = GetToken(c)
	})
}

// 设置客户端
func SetupClient() {
	client := &http.Client{Timeout: 5 * time.Second}
	//请求授权码
	err := SendPOST(client, GetTokenURL(AuthServerInfo.AuthURL, *clientconfig))
	if err != nil {
		log.Println("FAIL TO GET TOKEN")
		return
	}

	//使用令牌向服务器发送请求

}

// 设置请求对象
func SendPOST(client *http.Client, authurl string) error {
	//新建请求对象
	req, err := http.NewRequest("POST", authurl, strings.NewReader(""))
	if err != nil {
		return nil
	}
	// 使用req.Header.Set方法设置请求头中的Content-Type为application/json，表示请求体是JSON格式的数据
	req.Header.Set("Content-Type", "application/json")

	// 使用client.Do方法发送请求，并获取响应对象
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	//解析json获取令牌
	MyToken, err = GetTokenFromJson(resp)
	if err != nil {
		return err
	}

	// 延迟关闭响应对象的Body字段，释放资源
	defer resp.Body.Close()
	return nil
}

// TODO:使用令牌向服务器发送请求
func SendGET(client *http.Client, targeturl string) error {
	req, err := http.NewRequest("POST", targeturl, strings.NewReader(""))
	if err != nil {
		return nil
	}
	// 使用req.Header.Set方法设置请求头中的Content-Type为application/json，表示请求体是JSON格式的数据
	req.Header.Set("Authorization", "token "+MyToken)

	// 使用client.Do方法发送请求
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// 延迟关闭响应对象的Body字段，释放资源
	defer resp.Body.Close()
	return nil

}
