package database

import (
	"OIDC/oauth2service/model"
	"OIDC/oauth2service/myerror"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error
	DB, err = gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/oauth2"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	SetTable(DB)
	return nil
}

func SetTable(Db *gorm.DB) {
	Db.AutoMigrate(&model.Token{}, &model.ClientInfo{}, &model.User{})
}

func QueryClientByID(ClientID string) (*model.ClientInfo, error) {
	var client model.ClientInfo
	err := DB.Where("client_id=?", ClientID).First(&client)
	if err != nil {
		return nil, myerror.ErrInvalidClientInfo
	}
	return &client, nil
}

func QuerytokenByCode(AccessCode string) (*model.Token, error) {
	var token model.Token
	err := DB.Where("code=?", AccessCode).First(&token).Error
	if err != nil {
		return nil, myerror.ErrInvalidAuthorizeCode
	}
	//同时根据clientid联合查询clientid和clientsecret
	return &token, err
}
func StoreClientInfo(client model.ClientInfo) error {
	err := DB.Create(&client).Error
	if err != nil {
		return myerror.ErrServerError
	}
	return nil
}

func StoreToken(token model.Token) error {
	err := DB.Create(&token).Error
	if err != nil {
		return myerror.ErrServerError
	}
	return nil
}

func UpdateToken(token model.Token) error {
	err := DB.Save(&token).Error
	if err != nil {
		return myerror.ErrServerError
	}
	return nil
}

// 查找client_id和secret是否正确
func CheckClient(client model.ClientInfo) bool {
	var clientInfo model.ClientInfo
	result := DB.Where("client_id = ? and client_secret = ?", client.ClientID, client.ClientSecret).First(&clientInfo)
	return result.RowsAffected > 0
}

// TODO: 查找token是否存在且匹配
// 如果匹配返回用户信息
func CheckToken(token string) *model.User {
	var user model.User
	result := DB.Joins("JOIN users ON token.user_id = user.id").Where("token.access = ?", token).First(&user)
	if result == nil {
		return nil
	}
	return &user
}

// 检查code,clientid和clientsecret是否匹配
// 如果匹配返回token，否则返回nil
func CheckAccessCode(a model.QueryTokenRequest) *model.Token {
	var token model.Token
	//根据accesscode查找client_id,client_secret是否正确
	err := DB.Joins("INNER JOIN client ON tokens.client_id = registered_clients.client_id").
		Where("tokens.code = ? AND registered_clients.client_id = ? AND registered_clients.client_secret = ?", a.Code, a.ClientID, a.ClientSecret).
		First(&token).Error
	if err != nil {
		return nil
	}
	return &token
}

// 检查用户id和密码是否匹配
func CheckUser(user model.User) bool {
	var queryuser model.User
	result := DB.Where("userid = ? and password = ?", user.UserID, user.Password).First(&queryuser)
	return result.RowsAffected > 0
}

// 检查refresh_token是否过期且存在
func CheckRefreshToken(refreshtoken string) *model.Token {
	var token model.Token
	result := DB.Where("refresh = ?", refreshtoken).First(&token)
	currentTime := time.Now()

	// 比较当前时间与刷新令牌的过期时间

	if result.RowsAffected == 0 || currentTime.After(token.RefreshExpiresIn) {
		return nil
	}
	return &token
}
