package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	err := r.Run(":5300")
	if err != nil {
		fmt.Println("服务启动失败:", err)
		return
	}

	fmt.Println("服务启动成功")
}
