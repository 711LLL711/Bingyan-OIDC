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
	"strings"

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
	userinfo, err := database.UserRegister(&user)
	if err != nil {
		utils.Logger.Info("error while registering user")
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}

	//TODO:发送邮件

	//生成verifytoken
	verificationCode := utils.GenerateVerificationCode()
	userinfo.VerificationToken = verificationCode
	//存入数据库
	err = database.UpdateVerified(userinfo)
	if err != nil {
		utils.Logger.Info("error while updating user")
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	var firstName = user.Username

	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}
	//生成邮件内容
	emaildata := utils.EmailData{
		URL:       "http://localhost:8080/verify?token=" + verificationCode,
		FirstName: firstName,
		Subject:   "Verify your email",
	}
	//发送邮件
	utils.SendEmail(&user, &emaildata)
	message := "A verification email has been sent to " + user.Email + "."
	c.JSON(http.StatusOK, response.MakeSucceedResponse(map[string]interface{}{"msg": message}))
}

func UserLogin(c *gin.Context) {
	var user request.UserLogInRequest

	// user.Email = c.PostForm("email")
	// user.Password = c.PostForm("password")
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	//数据库查询,验证密码+是否已经验证邮箱

	resp, err := database.UserLogin(user)
	if err != nil {
		utils.Logger.Info("error while logging in")
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	//生成token

	token := utils.GenerateJWTToken(strconv.FormatUint(uint64((*resp).UserID), 10))
	//登陆成功，返回结果
	c.JSON(http.StatusOK, response.MakeSucceedResponse(map[string]interface{}{"user": *resp, "token": token}))
}

// 支持更新用户名和头像
func UserUpdate(c *gin.Context) {
	var user model.User
	//从token中获取用户id
	user.ID = c.MustGet("userID").(int)
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

		err = c.SaveUploadedFile(file, imgurl)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.InternalServerError)
			return
		}
		user.Avatar = "localhost:8080/" + imgurl
	}
	resp, err := database.UserUpdate(&user)
	fmt.Println("resp:", resp)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.DatabaseError)
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(*resp))
}

// TODO:增加验证邮箱的路由
// verify?token=xxxx
func UserVerify(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	//查询数据库
	err := database.VerifyUser(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.UnVerifiedError)
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(nil))
}
