package main

import "github.com/gin-gonic/gin"

func pong(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "pong",
	})
}
func main() {
	//实例化
	r := gin.Default()
	r.GET("/ping", pong)
	r.Run() // listen and serve on 0.0.0.0:8080
}
