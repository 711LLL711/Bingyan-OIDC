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
	"github.com/go-playground/validator/v10"
)

// 用户注册请求，输入邮箱密码注册
func UserRegister(c *gin.Context) {
	var user request.UserRegisterRequest
	// user.Email = c.PostForm("email")
	// user.Username = c.PostForm("username")
	// user.Password = c.PostForm("password")
	if err := c.ShouldBind(&user); err != nil {
		utils.Logger.Error("error while binding json")
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}

	//检查数据合法性
	if err := validator.New().Struct(user); err != nil {
		utils.Logger.Info("error while validating user")
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}

	//插入数据库
	reponse, err := database.UserRegister(&user)
	if err != nil {
		utils.Logger.Info("error while registering user")
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}

	//TODO:发送邮件

	//注册成功
	c.JSON(http.StatusOK, response.MakeSucceedResponse(reponse))
}

func UserLogin(c *gin.Context) {
	var user request.UserLogInRequest

	// user.Email = c.PostForm("email")
	// user.Password = c.PostForm("password")
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	//数据库查询

	resp, err := database.UserLogin(user)
	if err != nil {
		utils.Logger.Info("error while logging in")
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	//生成token
	fmt.Println(*resp)
	fmt.Println("USER LOGIN函数中user id:", (*resp).UserID)
	token := utils.GenerateJWTToken(strconv.FormatUint(uint64((*resp).UserID), 10))
	//登陆成功，返回结果
	c.JSON(http.StatusOK, response.MakeSucceedResponse(map[string]interface{}{"user": *resp, "token": token})) //这里要解引用！
}

// 支持更新用户名和头像
func UserUpdate(c *gin.Context) {
	var user model.User
	//从token中获取用户id
	user.ID = c.MustGet("userID").(uint)
	//log.Println("user id:", user.ID)
	user.Username = c.PostForm("username")
	user.Bio = c.PostForm("bio")
	file, _ := c.FormFile("avatar")
	if file != nil {
		//生成头像url
		imgurl, err := utils.GenerateImgName(file.Filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.InvalidInfoError)
			return
		}
		c.SaveUploadedFile(file, imgurl)
		user.Avatar = "locolhost:8080/" + imgurl
	}
	resp, err := database.UserUpdate(&user)
	fmt.Println("resp:", resp)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.DatabaseError)
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(*resp))
}
