package main

import (
	"github.com/gin-gonic/gin"
	"go-egg/internal/server"
)

// main todo 后期需要迁移到 cmd 目录下面
func main() {
	router := gin.New()

	server.SetupRoute(router)

	// 运行服务
	err := router.Run(":7001")
	if err != nil {
		panic(err.Error())
	}
}
