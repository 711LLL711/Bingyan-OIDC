package server

import (
	"OIDC/oauth2service/database"
	"OIDC/oauth2service/manage"
	"OIDC/oauth2service/model"
	"OIDC/oauth2service/myerror"
	"OIDC/oauth2service/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	c := gin.Default()
	//注册
	c.GET("/appication", func(c *gin.Context) {
		c.HTML(http.StatusOK, "application.html", nil)
	})
	c.POST("/appication", HandleClientRegister)

	//授权码模式处理用户发送的授权码请求
	// /authorize?response_type=code&client_id=123456&redirect_uri=http://localhost:9094/oauth2&scope=scope1
	c.GET("/authorize", HandleAuthorizeCodeRequest)
	//token用来签发token
}

// 从url中解析客户端的client_id和client_secret
func GetClientIDAndSecret(c *gin.Context) (client_id string, client_secret string, err error) {
	client_id = c.Query("client_id")
	client_secret = c.Query("client_secret")
	if client_id == "" || client_secret == "" {
		return "", "", myerror.ErrInvalidClientInfo
	}
	return client_id, client_secret, nil
}

// 从url中解析授权码
func GetAccessCode(c *gin.Context) (code string, err error) {
	code = c.Query("code")
	if code == "" {
		return "", myerror.ErrInvalidAuthorizeCode
	}
	return code, nil
}

// 从请求中解析token
func GetToken(c *gin.Context) (token string, err error) {
	//Authorization头部解析token
	auth := c.GetHeader("Authorization")

	if len(auth) <= len("Token ") {
		return "", myerror.ErrInvalidAccessToken
	}
	tokenString := auth[len("Token "):]
	return tokenString, nil
}

// 处理并存储客户端注册信息，颁发client_id和client_secret
func HandleClientRegister(c *gin.Context) {
	var clientInfo model.ClientInfo
	clientInfo.ClientName = c.PostForm("ClientName")
	clientInfo.ClientDomain = c.PostForm("domain")
	clientInfo.RedirectURI = c.PostForm("redirect_uri")
	clientInfo.ClientID, clientInfo.ClientSecret = manage.GenerateClientIDAndSecret()
	err := database.StoreClientInfo(clientInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.MakeFailedResponseBody(myerror.Descriptions[err]))
	}
	c.JSON(http.StatusOK, response.MakeResponseBody("", map[string]interface{}{"client_id": clientInfo.ClientID, "client_secret": clientInfo.ClientSecret}))
}

func HandleAccessToken(c *gin.Context) {
	AccessCode, err := GetAccessCode(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponseBody(myerror.Descriptions[err]))
		return
	}
	//查询是否正确
	token, err := database.QueryClientByCode(AccessCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponseBody(myerror.Descriptions[err]))
	}

	//TODO:检查是否过期
}

func HandleAuthorizeCodeRequest(c *gin.Context)
