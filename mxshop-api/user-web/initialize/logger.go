package initialize

import (
	"fmt"
	"go.uber.org/zap"
)

func InitLogger() {
	fmt.Println("【InitLogger】初始化")
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
