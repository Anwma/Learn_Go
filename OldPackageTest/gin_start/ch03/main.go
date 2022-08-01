package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//
func main() {
	router := gin.Default()
	//路由分组
	goodsGroup := router.Group("/goods")
	{
		//下面两行产生冲突,导致异常
		//panic: wildcard segment ':id' conflicts with existing children in path '/goods/: id'
		goodsGroup.GET("", goodsList)
		//goodsGroup.GET("/:id/:action", goodsDetail) //获取商品id为1的详细信息 符合某种模式
		goodsGroup.GET("/:id/*action", goodsDetail) //{"action":"/delete/hello/a",
		goodsGroup.POST("", createGoods)
	}

	router.Run(":8082")
}

func createGoods(c *gin.Context) {

}

func goodsDetail(c *gin.Context) {
	id := c.Param("id")
	action := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"action": action,
	})
}

func goodsList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"name": "goodsList",
	})
}
