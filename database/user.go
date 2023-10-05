package database

import (
	"OIDC/model"
	"OIDC/model/request"
	"OIDC/model/response"
	"OIDC/utils"
	"errors"

	"github.com/jinzhu/copier"
)

// UserRegister 用户注册 - 调用/查重/函数和/创建/用户函数
func UserRegister(userReq *request.UserRegisterRequest) (*response.UserResponse, error) {
	//检查邮箱是否存在
	// 检查邮箱是否存在
	var user model.User
	result := DB.Where("email = ?", userReq.Email).First(&user)
	if result.RowsAffected > 0 {
		return nil, errors.New("already exist")
	}

	// 创建用户
	userResponse, err := createUser(userReq)
	if err != nil {
		utils.Logger.Error("创建用户错误\n")
		return nil, err
	}
	return userResponse, nil

}

// createUser 创建用户 - 内部函数
func createUser(userReq *request.UserRegisterRequest) (*response.UserResponse, error) {
	// 改用 copier 进行结构体转换
	user := model.User{}
	err := copier.Copy(&user, userReq)
	if err != nil {
		utils.Logger.Error("用户注册时,结构体转换错误\n")
		return nil, err
	}
	// 密码加密
	err = utils.PasswordHash(&user.Password)
	if err != nil {
		utils.Logger.Error("密码加密错误")
		return nil, err
	}

	// //生成用户id，改为自增id
	// user.ID= utils.GenerateRandomID(8)
	// 存入数据库
	if err := DB.Create(&user).Error; err != nil {
		utils.Logger.Error("创建用户错误\n")
		return nil, err
	}
	// TODO:生成用户头像

	// 赋予用户名和头像
	// width := 8
	// user.UserName = fmt.Sprintf("User%0*s", width, strconv.FormatInt(int64(user.UserID), 10))
	// rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	// rndNum := rng.Intn(30)
	// user.UserAvatar = fmt.Sprintf(global.Server.Avatar.UserUrl, rndNum)
	// // 进行更新
	// if err := global.Sql.Model(&user).Updates(user).Error; err != nil {
	// 	global.Logger.Errorf("更新用户名和头像错误\n")
	// 	return nil, err
	// }
	// 进行返回
	var userResponse response.UserResponse
	err = copier.Copy(&userResponse, &user)
	if err != nil {
		utils.Logger.Error("用户注册时,结构体转换错误\n")
		return nil, err
	}
	return &userResponse, nil
}

func UserLogin(userReq request.UserLogInRequest) (*response.UserResponse, error) {
	//查询邮箱是否存在
	var user model.User
	// if result := DB.Where("email = ?", userReq.Email).First(&user); result.RowsAffected == 0 {
	// 	utils.Logger.Info("邮箱不存在")
	// 	return nil, result.Error
	// }

	if err := DB.Where("email = ?", userReq.Email).First(&user).Error; err != nil {
		utils.Logger.Info("查询错误")
		return nil, err
	}
	//fmt.Println("user:", user)
	//验证密码
	if !utils.PasswordVerify(user.Password, userReq.Password) {
		utils.Logger.Info("incorrect password")
		return nil, errors.New("incorrect password")
	}

	//登陆成功
	var userResponse response.UserResponse
	//copier.Copy(&userResponse, &user)
	userResponse.UserID = user.ID
	userResponse.UserName = user.Username
	userResponse.UserAvatar = user.Avatar
	return &userResponse, nil
}

func UserUpdate(user *model.User) (*response.UserResponse, error) {
	var userResponse response.UserResponse
	if user.Password != "" {
		utils.PasswordHash(&user.Password)
	}
	if err := DB.Model(&user).Updates(user).Where("id = ?", user.ID).Error; err != nil {
		utils.Logger.Error("更新用户信息错误\n")
		return nil, err
	}
	//copier.Copy(&userResponse, &user)
	DB.Where("id = ?", user.ID).First(&user)
	userResponse.UserID = user.ID
	userResponse.UserName = user.Username
	userResponse.UserAvatar = user.Avatar
	return &userResponse, nil
}
