package main

import (
	"fmt"
	"go.uber.org/zap"
	"mxshop-api/user-web/initialize"
)

func main() {
	port := 8021

	//1.初始化logger
	initialize.InitLogger()

	//2.初始化routers
	Router := initialize.Routers()


	/*
		1.S()可以获取一个全局的sugar 可以让我们自己设置一个全局的logger
		2.日志是分级别的 debug info warn error fatal error
		3.S 函数和 L 函数很有用, 提供了一个全局的安全访问logger的途径
	*/
	//logger, _ := zap.NewProduction()
	//defer logger.Sync()
	//suger := logger.Sugar()

	zap.S().Debugf("启动服务器,端口：%d", port)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}

}
