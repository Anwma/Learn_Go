package plugin

import "github.com/gin-gonic/gin"

const OnlyFuncName = "Plugin"

type Plugin interface {
	// Register 注册路由
	Register(group *gin.RouterGroup)

	//RouterPath 用户返回注册路由
	RouterPath() string
}
