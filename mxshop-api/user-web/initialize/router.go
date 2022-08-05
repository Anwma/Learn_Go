package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/router"
)

func Routers() *gin.Engine {
	fmt.Println("【Routers】初始化")
	Router := gin.Default()

	ApiGroup := Router.Group("/u/v1")
	router.InitUserRouter(ApiGroup)

	return Router
}
