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

func InitRouter() *gin.Engine {
	c := gin.Default()
	c.LoadHTMLGlob("templates/*")
	//客户端注册
	c.GET("/appication", func(c *gin.Context) {
		c.HTML(http.StatusOK, "application.html", nil)
	})
	//接收客户端注册信息,赋予client_id和client_secret
	c.POST("/application", HandleClientRegister)

	//授权码模式处理用户发送的授权码请求
	// /authorize?response_type=code&client_id=123456&redirect_uri=http://localhost:9094/oauth2&scope=scope1

	//请求授权码，服务器要求用户登录，确认用户信息
	c.GET("/authorize", HandleAuthorizeCodeRequest)
	//处理用户登录信息确认授权
	c.POST("/auth/login", HandleAuthLogin)

	//签发token
	c.POST("/auth/token", HandleTokenRequest)

	//处理用token请求数据的路由
	c.GET("api/user", HandleUserInfoRequest)
	return c
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
	clientInfo.ClientName = c.PostForm("clientname")
	clientInfo.ClientDomain = c.PostForm("domain")
	clientInfo.RedirectURI = c.PostForm("redirect_url")
	clientInfo.ClientID, clientInfo.ClientSecret = manage.GenerateClientIDAndSecret()
	//存储客户端信息
	err := database.StoreClientInfo(clientInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.MakeFailedResponseBody(myerror.Descriptions[err]))
	}
	c.JSON(http.StatusOK, response.MakeResponseBody("", map[string]interface{}{"client_id": clientInfo.ClientID, "client_secret": clientInfo.ClientSecret}))
}

// 处理客户端请求授权码
// 客户端传入client_id和稍后重定向回的redirect_uri
// 服务器要求用户登录，如果登陆成功并确认授权
// 服务器把授权码code传进url参数中，重定向回redirect_uri
// 如果授权失败，告诉resource owner 授权失败,返回error=access_denied
func HandleAuthorizeCodeRequest(c *gin.Context) {
	//解析 clientid和redirect_uri
	clientid := c.Query("client_id")
	redirect_uri := c.Query("redirect_uri")
	c.Redirect(http.StatusFound, "/auth/login?clientid="+clientid+"&redirect_uri="+redirect_uri)
}

// 处理用户登录信息确认授权
func HandleAuthLogin(c *gin.Context) {
	clientid := c.Query("client_id")
	redirect_uri := c.Query("redirect_uri")
	var userinfo model.User
	userinfo.UserID = c.PostForm("userid")
	userinfo.Password = c.PostForm("password")
	if !database.CheckUser(userinfo) { //检查用户是否存在且密码正确
		c.JSON(http.StatusUnauthorized, response.MakeFailedResponseBody(myerror.Descriptions[myerror.ErrAccessDenied]))
	}
	//生成授权码
	var token model.Token
	token.ClientID = clientid
	token.UserID = userinfo.UserID
	token.RedirectURI = redirect_uri
	token.Code = manage.GenerateAhthorizationCode()
	//存储授权码
	err := database.StoreToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.MakeFailedResponseBody(myerror.Descriptions[err]))
	}

	//重定向回redirect_uri并且带上code参数
	c.Redirect(http.StatusFound, redirect_uri+"?code="+token.Code)
}

// 如果客户端提供的autorizationcode有效，服务器签发token令牌
// 返回，userid,access_token,token_type,expires_in,refresh_token,scope
func HandleTokenRequest(c *gin.Context) {
	//解析出client_id和client_secret，authorizationcode
	var queryinfo model.QueryTokenRequest
	err := c.ShouldBind(&queryinfo)
	if err != nil || queryinfo.ClientID == "" || queryinfo.ClientSecret == "" || queryinfo.Code == "" {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponseBody(myerror.Descriptions[myerror.ErrInvalidAccessToken]))
		return
	}
	//检查三者是否正确且匹配,正确则获得token结构体指针
	token := database.CheckAccessCode(queryinfo)
	if token == nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponseBody(myerror.Descriptions[myerror.ErrInvalidAccessToken]))
		return
	}

	//TODO:检查token是否过期

	//正确则生成token
	GenerateToken := manage.GenerateToken(token.UserID)
	token.Access = GenerateToken
	//更新token
	database.UpdateToken(*token)
	//json数据AccessToken:xxxx
	//返回json数据
	c.JSON(http.StatusOK, response.MakeResponseBody("", map[string]interface{}{"access_token": GenerateToken, "token_type": "Bearer"}))
}

func HandleUserInfoRequest(c *gin.Context) {
	//解析token
	tokenString, err := GetToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.MakeFailedResponseBody(myerror.Descriptions[myerror.ErrAccessDenied]))
		return
	}
	//检查token是否有效
	user := database.CheckToken(tokenString)
	if user == nil {
		c.JSON(http.StatusForbidden, response.MakeFailedResponseBody(myerror.Descriptions[myerror.ErrInvalidAccessToken]))
		return
	}

	var userinfo = make(map[string]interface{})
	userinfo["userid"] = user.UserID
	userinfo["username"] = user.Username

	//返回用户信息
	c.JSON(http.StatusOK, response.MakeResponseBody("", userinfo))
}
