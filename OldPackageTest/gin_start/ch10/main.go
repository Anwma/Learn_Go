package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//优雅退出 当我们关闭程序的时候应该做的后续处理
	//微服务 启动之前 或者启动之后 会做一件事：将当前的服务器ip地址和端口号注册到注册中心
	//我们当前的服务停止了以后并没有告知注册中心
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"pong",
		})
	})
	go func() {
		router.Run(":8084")
	}()
	//如果想要接收到信号 kill -9 强杀命令
	quit := make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	//处理后续逻辑
	fmt.Println("关闭server中......")
	fmt.Println("注销服务")
}
