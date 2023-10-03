package utils

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	var err error
	Logger, err = zap.NewProduction() // NewProduction() 和 NewDevelopment() NewExample() 三种模式
	if err != nil {
		// 处理错误
		panic(err)
	}
	defer Logger.Sync() // 在程序退出时释放资源
}
