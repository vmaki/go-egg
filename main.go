package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	s := gin.Default()

	s.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Hello World!",
		})
	})

	// 运行服务
	s.Run(":7001")
}
