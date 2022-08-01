package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//
func main() {
	router := gin.Default()
	//路由分组
	goodsGroup := router.Group("/goods/1")
	{
		goodsGroup.GET("/list", goodsList)
		goodsGroup.GET("/1", goodsDetail) //获取商品id为1的详细信息 模式
		goodsGroup.POST("/add", createGoods)
	}



	router.Run(":8082")
}

func createGoods(context *gin.Context) {

}

func goodsDetail(context *gin.Context) {

}

func goodsList(context *gin.Context) {
	context.JSON(http.StatusOK,gin.H{
		"name":"goodsList",
	})
}
