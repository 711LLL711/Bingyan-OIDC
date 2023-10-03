package database

import (
	"OIDC/config"
	"OIDC/model"
	"OIDC/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn, err := config.SetDbConfig()
	if err != nil {
		utils.Logger.Error("error while loading config")
		return err
	}
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Logger.Error("error while connecting to database:")
		return err
	}

	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		utils.Logger.Error("error while migrating database:")
		return err
	}

	return nil
}

//不用viper读取配置文件
// func Connect() {
// 	var err error
// 	DB, err = gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/OIDC"), &gorm.Config{})
// 	if err != nil {
// 		fmt.Printf("连接失败")
// 		panic(err)
// 	}
// 	DB.AutoMigrate(&model.User{})
// }
