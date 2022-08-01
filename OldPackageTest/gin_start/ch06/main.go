package main

import (
	"OldPackageTest/gin_start/ch06/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/moreJSON", moreJSON)
	router.GET("/someProtoBuf", returnProto)

	router.Run(":8083")
}

func returnProto(c *gin.Context) {
	course := []string{"python", "go", "微服务"}
	user := &proto.Teacher{
		Name:   "bobby",
		Course: course,
	}
	c.ProtoBuf(http.StatusOK, user)
}

func moreJSON(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "jeff"
	msg.Message = "这是一个测试json"
	msg.Number = 20

	c.JSON(http.StatusOK, msg)

}
