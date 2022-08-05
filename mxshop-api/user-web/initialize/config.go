package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mxshop-api/user-web/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
	//我们刚才设置的环境变量 想要生效 必须得重启Goland
}
func InitConfig() {
	fmt.Println("【InitConfig】初始化")
	debug := GetEnvInfo("MXSHOP_DEBUG")
	//自定义前缀
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user-web/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("user-web/%s-debug.yaml", configFilePrefix)
	}
	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//这个对象如何在其他文件中使用 - 全局变量
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息: %v", global.ServerConfig)
	//The system cannot find the file specified.

	//viper的功能 - 动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		zap.S().Infof("配置文件产生变化: %s", in.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
		zap.S().Infof("配置信息: %v", global.ServerConfig)
	})
}
