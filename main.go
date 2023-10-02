package main

import (
	"OIDC/database"
	"OIDC/routes"
	"OIDC/utils"
)

func main() {
	utils.InitLogger()

	//TODO:viper读取配置文件出错，导致数据库未连接
	database.Connect()

	//创建路由引擎
	e := routes.InitRouter()
	//启动服务
	e.Run(":8080")
}
