package main

import (
	"OIDC/database"
	"OIDC/routes"
	"OIDC/utils"
)

func main() {
	utils.InitLogger()

	err := database.Connect()
	if database.DB == nil || err != nil {
		utils.Logger.Error("数据库连接失败")
		return
	}
	//创建路由引擎
	e := routes.InitRouter()
	//启动服务
	e.Run(":8080")
}
