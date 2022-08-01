package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   int `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}
//参数无法获取到
func main() {
	router := gin.Default()
	router.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		//err1:ShouldBindUri写成 ShouldBind
		if err := c.ShouldBindUri(&person); err != nil {
			//c.JSON(404, gin.H{"msg": err})
			c.Status(404)
			return
		}
		c.JSON(200, gin.H{
			"name": person.Name,
			"uuid": person.ID,
		})
	})
	router.Run(":8083")
}
