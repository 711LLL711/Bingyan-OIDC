package oauth2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 生成链接跳转到 授权服务器
// AuthUrl是授权服务器的地址
func GetAuthCodeURL(ServerURL string) string {
	//exp: http://localhost:8081/oauth/authorize
	//?client_id=ed9d28ab310cfce4c521
	//&redirect_uri=http://localhost:8080/oauth/redirect
	return ServerURL + "?client_id=" + clientconfig.ClientID + "&redirect_uri=" + clientconfig.RedirectURL
}

// 处理重定向回来的，从认证服务器获取授权码
//
//exp:http://localhost:8080/oauth/redirect?code=0c2d9b7b4b9c1b0b9b7b
func GetToken(c *gin.Context) string {
	return c.Query("code")
}

// 用授权码请求token令牌
// 生成请求令牌的url
// 头部： accept: 'application/json'
func GetTokenURL(ServerURL string, clientconfig Clientcode) string {
	return ServerURL + "?client_id=" + clientconfig.ClientID + "&client_secret=" + clientconfig.ClientSecret + "&grant_type=authorization_code&code=" + clientconfig.Code
}

// 从服务器发来的响应体解析token令牌
func GetTokenFromJson(resp *http.Response) (string, error) {
	// 创建一个 map 来存储 JSON 数据
	var data map[string]interface{}

	// 读取响应体的 JSON 数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 使用 json.Unmarshal 函数解码 JSON 数据到 map
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", nil
	}

	// 现在，你可以根据键来获取特定字段的值
	codeValue, ok := data["code"].(string)
	if !ok {
		fmt.Println("Field 'code' not found or not a number")
		return "", fmt.Errorf("Field 'code' not found or not a number")
	}

	// 打印 'code' 字段的值
	fmt.Println("Code:", string(codeValue))
	return string(codeValue), nil
}

//生成包含token的请求头
