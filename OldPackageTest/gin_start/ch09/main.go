package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	router := gin.Default()
	//LoadHTMLFiles 会将指定目录下的文件加载好，相对目录
	//为什么我们通过golang运行main.go的时候并没有生成main.exe文件
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(dir)
	//C:\Users\Anwma\AppData\Local\Temp\GoLand
	//推荐在指定目录下 go build 后运行
	router.LoadHTMLGlob("templates/**/*")
	//router.LoadHTMLFiles("templates/index.tmpl", "templates/goods.html")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Jetbrains",
		})
	})

	router.GET("/goods/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", gin.H{
			"name": "Golang",
		})
	})

	router.GET("/users/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", gin.H{
			"name": "Golang",
		})
	})

	router.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"name": "Golang",
		})
	})

	router.Run(":8083")

}
