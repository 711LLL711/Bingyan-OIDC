package controller

import (
	"OIDC/database"
	"OIDC/model"
	"OIDC/model/request"
	"OIDC/model/response"
	"OIDC/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 用户注册请求，输入邮箱密码注册
func UserRegister(c *gin.Context) {
	var user request.UserRegisterRequest
	user.Email = c.PostForm("email")
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	utils.Logger.Error("error while binding json")
	// 	c.JSON(http.StatusBadRequest, response.InvalidInfoError)
	// 	return
	// }

	//检查数据合法性
	// if err := validator.New().Struct(user); err != nil {
	// 	utils.Logger.Error("error while validating user")
	// 	c.JSON(http.StatusBadRequest, response.InvalidInfoError)
	// 	return
	// }

	//插入数据库
	reponse, err := database.UserRegister(&user)
	if err != nil {
		utils.Logger.Error("error while registering user")
		c.JSON(http.StatusBadRequest, response.DatabaseError)
		return
	}

	//发送邮件

	//注册成功
	c.JSON(http.StatusOK, response.MakeSucceedResponse(reponse))
}

func UserLogin(c *gin.Context) {
	var user request.UserLogInRequest

	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, response.InvalidInfoError)
	// 	return
	// }

	//TODO:这一步正常
	fmt.Println("user:", user)
	//数据库查询

	resp, err := database.UserLogin(user)
	if err != nil {
		utils.Logger.Error("error while logging in")
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	//生成token
	token := utils.GenerateJWTToken(strconv.FormatUint(uint64(resp.UserID), 10))
	//登陆成功，返回结果
	c.JSON(http.StatusOK, response.MakeSucceedResponse(map[string]interface{}{"user": *resp, "token": token})) //这里要解引用！
}

func UserUpdate(c *gin.Context) {
	var user model.User
	resp, err := database.UserUpdate(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.DatabaseError)
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(resp))
}
