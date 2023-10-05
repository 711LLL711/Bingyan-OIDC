package database

import (
	"OIDC/oauth2service/model"
	"OIDC/oauth2service/myerror"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error
	DB, err = gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/oauth2"), &gorm.Config{})
	if err != nil {
		return err
	}
	SetTable(DB)
	return nil
}

func SetTable(Db *gorm.DB) {
	Db.AutoMigrate(&model.Token{}, &model.RegisteredClient{})
}

func QueryClientByID(ClientID string) (model.RegisteredClient, error) {
	var client model.RegisteredClient
	err := DB.Where("client_id=?", ClientID).First(&client).Error
	return client, err
}

func QueryClientByCode(AccessCode string) (*model.RegisteredClient, error) {
	var client model.RegisteredClient
	err := DB.Where("access_code=?", AccessCode).First(&client).Error
	if err != nil {
		return nil, myerror.ErrInvalidAuthorizeCode
	}
	return &client, err
}
func StoreClientInfo(client model.ClientInfo) error {
	err := DB.Create(&client).Error
	return err
}

func StoreToken(token model.Token) error {
	err := DB.Create(&token).Error
	if err != nil {
		return myerror.ErrServerError
	}
	return nil
}

// 查找client_id和secret是否正确
func CheckClient(client model.ClientInfo) bool {
	var clientInfo model.ClientInfo
	result := DB.Where("client_id = ? and client_secret = ?", client.GetID(), client.GetSecret()).First(&clientInfo)
	return result.RowsAffected > 0
}

// 查找accessCode是否正确
func CheckAccessCode(token model.Token) bool {
	var tokenInfo model.Token
	result := DB.Where("client_id = ?", token.ClientID).First(&tokenInfo)
	if result.RowsAffected == 0 {
		return false
	}
	return tokenInfo.Access == token.Access

}

// TODO: 查找token是否存在且匹配
func CheckToken(token model.Token) {
	return
}
