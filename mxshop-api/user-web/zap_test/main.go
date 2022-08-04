package main

import (
	"go.uber.org/zap"
)

func main() {
	//logger, _ := zap.NewProduction() //生产环境
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	//url := "https://imooc.com"
	logger.Info("failed to fetch URL")
	//生产环境： {"level":"info","ts":1659612291.9787142,"caller":"zap_test/main.go:12","msg":"failed to fetch URL"}
	//开发环境：2022-08-04T19:25:17.586+0800    INFO    zap_test/main.go:12     failed to fetch URL
	//sugar := logger.Sugar()
	//sugar.Infow("failed to fetch URL",
	//	// Structured context as loosely typed key-value pairs.
	//	"url", url,
	//	"attempt", 3,
	//)
	//sugar.Infof("Failed to fetch URL: %s", url)
}
