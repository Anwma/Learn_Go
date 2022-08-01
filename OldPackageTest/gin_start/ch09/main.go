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
	router.LoadHTMLFiles("templates/index.tmpl")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Jetbrains",
		})
	})

	router.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"name": "Golang",
		})
	})

	router.Run(":8083")

}
